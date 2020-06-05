package control

import (
	"io"
	"sync"

	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
)

type term struct {
	prefix string
	p      execution.P
	err    error

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
	redirector := &iostream.StdWriters{
		Stdout: iostream.NewTerminal(p.prefix, p.outW),
		Stderr: iostream.NewTerminal(p.prefix, p.errW),
	}
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
		ioDone := results.Stream(redirector)
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

func Term(prefix string, q execution.P) execution.P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &term{
		prefix: prefix,
		p:      q,
		outR:   outR,
		outW:   outW,
		errR:   errR,
		errW:   errW,
	}
	p.wg.Add(1)
	return p
}
