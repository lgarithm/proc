package control

import (
	"io"
	"sync"

	"github.com/lgarithm/proc/iostream"
)

type term struct {
	redirector *iostream.StdWriters
	p          P
	err        error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *term) Stdpipe() (io.Reader, io.Reader, error) {
	return p.outR, p.errR, nil
}

func (p *term) Start() error {
	go func() {
		defer p.wg.Done()
		defer p.errW.Close()
		defer p.outW.Close()
		stdout, stderr, err := p.p.Stdpipe()
		if err != nil {
			p.err = err
			return
		}
		results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
		ioDone := results.Stream(p.redirector)
		if err := p.p.Start(); err != nil {
			p.err = err
			return
		}
		ioDone.Wait()
		if err := p.p.Wait(); err != nil {
			p.err = err
			return
		}
	}()
	return nil
}

func (p *term) Wait() error {
	p.wg.Wait()
	return p.err
}

func Term(prefix string, q P) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	ps := iostream.PromStr(prefix)
	redirector := &iostream.StdWriters{
		Stdout: ps.NewTerm(outW),
		Stderr: ps.NewTerm(errW),
	}
	p := &term{
		redirector: redirector,
		p:          q,
		outR:       outR,
		outW:       outW,
		errR:       errR,
		errW:       errW,
	}
	p.wg.Add(1)
	return p
}

func DynTerm(ps iostream.PromoFn, q P) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	redirector := &iostream.StdWriters{
		Stdout: ps.NewTerm(outW),
		Stderr: ps.NewTerm(errW),
	}
	p := &term{
		redirector: redirector,
		p:          q,
		outR:       outR,
		outW:       outW,
		errR:       errR,
		errW:       errW,
	}
	p.wg.Add(1)
	return p
}
