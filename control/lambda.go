package control

import (
	"io"
	"sync"
)

type lambda struct {
	p   func() P
	err error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *lambda) Stdpipe() (io.Reader, io.Reader, error) {
	return p.outR, p.errR, nil
}

func (p *lambda) Start() error {
	redirector := &StdWriters{
		Stdout: p.outW,
		Stderr: p.errW,
	}
	go func() {
		defer p.wg.Done()
		defer p.errW.Close()
		defer p.outW.Close()
		q := p.p()
		stdout, stderr, err := q.Stdpipe()
		if err != nil {
			p.err = err
			return
		}
		results := StdReaders{Stdout: stdout, Stderr: stderr}
		ioDone := results.Stream(redirector)
		if err := q.Start(); err != nil {
			p.err = err
			return
		}
		ioDone.Wait()
		if err := q.Wait(); err != nil {
			p.err = err
			return
		}
	}()
	return nil
}

func (p *lambda) Wait() error {
	p.wg.Wait()
	return p.err
}

func Lambda(q func() P) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &lambda{
		p:    q,
		outR: outR,
		outW: outW,
		errR: errR,
		errW: errW,
	}
	p.wg.Add(1)
	return p
}
