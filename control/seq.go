package control

import (
	"io"
	"sync"

	"github.com/lgarithm/proc-experimental/iostream"
)

type seq struct {
	ps       []P
	firstErr error

	outR io.ReadCloser
	outW io.WriteCloser
	errR io.ReadCloser
	errW io.WriteCloser

	wg sync.WaitGroup
}

func (p *seq) Stdpipe() (io.Reader, io.Reader, error) {
	return p.outR, p.errR, nil
}

func (p *seq) Start() error {
	redirector := &iostream.StdWriters{
		Stdout: p.outW,
		Stderr: p.errW,
	}
	go func() {
		for _, q := range p.ps {
			stdout, stderr, err := q.Stdpipe()
			if err != nil {
				p.firstErr = err
				break
			}
			results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
			ioDone := results.Stream(redirector)
			if err := q.Start(); err != nil {
				p.firstErr = err
				break
			}
			ioDone.Wait()
			if err := q.Wait(); err != nil {
				p.firstErr = err
				break
			}
		}
		p.outW.Close()
		p.errW.Close()
		p.wg.Done()
	}()
	return nil
}

func (p *seq) Wait() error {
	p.wg.Wait()
	return p.firstErr
}

func Seq(ps ...P) P {
	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	p := &seq{
		ps:   ps,
		outR: outR,
		outW: outW,
		errR: errR,
		errW: errW,
	}
	p.wg.Add(1)
	return p
}
