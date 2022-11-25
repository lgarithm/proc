package control

import (
	"io"
	"sync"

	"github.com/lgarithm/proc/iostream"
)

type try struct {
	p       func(int) P
	lastErr error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *try) Stdpipe() (io.Reader, io.Reader, error) {
	return p.outR, p.errR, nil
}

func (p *try) Start() error {
	redirector := &iostream.StdWriters{
		Stdout: p.outW,
		Stderr: p.errW,
	}
	go func() {
		for i := 0; ; i++ {
			q := p.p(i)
			stdout, stderr, err := q.Stdpipe()
			if err != nil {
				p.lastErr = err
				continue
			}
			results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
			ioDone := results.Stream(redirector)
			if err := q.Start(); err != nil {
				p.lastErr = err
				continue
			}
			ioDone.Wait()
			if err := q.Wait(); err != nil {
				p.lastErr = err
				continue
			}
			p.lastErr = nil
			break
		}
		p.outW.Close()
		p.errW.Close()
		p.wg.Done()
	}()
	return nil
}

func (p *try) Wait() error {
	p.wg.Wait()
	return p.lastErr
}

func TryI(q func(int) P) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &try{
		p:    q,
		outR: outR,
		outW: outW,
		errR: errR,
		errW: errW,
	}
	p.wg.Add(1)
	return p
}

func Try(q func() P) P {
	return TryI(func(int) P { return q() })
}
