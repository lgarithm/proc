package control

import (
	"io"
	"sync"
)

type tee struct {
	p   P
	ws  []*StdWriters
	err error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *tee) Stdpipe() (io.Reader, io.Reader, error) {
	return p.outR, p.errR, nil
}

func (p *tee) Start() error {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer p.errW.Close()
		defer p.outW.Close()
		stdout, stderr, err := p.p.Stdpipe()
		if err != nil {
			p.err = err
			return
		}
		upstream := StdReaders{Stdout: stdout, Stderr: stderr}
		w0 := &StdWriters{
			Stdout: p.outW,
			Stderr: p.errW,
		}
		ws := append(p.ws, w0)
		ioDone := upstream.Stream(ws...)
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

func (p *tee) Wait() error {
	p.wg.Wait()
	return p.err
}

func Tee(q P, ws ...*StdWriters) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &tee{
		p:    q,
		ws:   ws,
		outR: outR,
		outW: outW,
		errR: errR,
		errW: errW,
	}
	return p
}
