package control

import (
	"io"
	"sync"

	"github.com/lgarithm/proc-experimental/iostream"
)

type par struct {
	ps []P

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
		out, err, e := q.Stdpipe()
		if e != nil {
			return nil, nil, e
		}
		outs = append(outs, out)
		errs = append(errs, err)
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
		go func(i int, q P) {
			q.Start()
			p.errs[i] = q.Wait()
			p.wg.Done()
		}(i, q)
	}
	return nil
}

func (p *par) Wait() error {
	p.wg.Wait()
	// FIXME: merge p.errs
	return nil
}

func Par(ps ...P) P {
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
