package control

import (
	"fmt"
	"io"
	"sync"

	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
)

type par struct {
	ps []execution.P

	errs []error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *par) Stdpipe() (io.Reader, io.Reader, error) {
	var outs, errs []io.Reader
	for _, q := range p.ps {
		stdout, stderr, err := q.Stdpipe()
		if err != nil {
			return nil, nil, err
		}
		outs = append(outs, stdout)
		errs = append(errs, stderr)
	}
	go func() {
		iostream.Mix(p.outW, outs...)
		p.outW.Close()
	}()
	go func() {
		iostream.Mix(p.errW, errs...)
		p.errW.Close()
	}()
	return p.outR, p.errR, nil
}

func (p *par) Start() error {
	for i, q := range p.ps {
		go func(i int, q execution.P) {
			q.Start()
			p.errs[i] = q.Wait()
			p.wg.Done()
		}(i, q)
	}
	return nil
}

func (p *par) Wait() error {
	p.wg.Wait()
	return mergeErrors(p.errs)
}

func Par(ps ...execution.P) execution.P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &par{
		ps:   ps,
		errs: make([]error, len(ps)),
		outR: outR,
		outW: outW,
		errR: errR,
		errW: errW,
	}
	p.wg.Add(len(ps))
	return p
}

func mergeErrors(errs []error) error {
	var msg string
	var failed int
	for _, e := range errs {
		if e != nil {
			failed++
			if len(msg) > 0 {
				msg += ", "
			}
			msg += e.Error()
		}
	}
	if failed == 0 {
		return nil
	}
	return fmt.Errorf("%d failed out of %d: %s", failed, len(errs), msg)
}
