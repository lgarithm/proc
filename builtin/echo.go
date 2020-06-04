package builtin

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type echo struct {
	s string

	outR io.ReadCloser
	outW io.WriteCloser

	wg sync.WaitGroup
}

func (p *echo) Stdpipe() (io.Reader, io.Reader, error) {
	err := &bytes.Buffer{}
	return p.outR, err, nil
}

func (p *echo) Start() error {
	fmt.Fprintf(p.outW, "%s", p.s)
	p.outW.Close()
	p.wg.Done()
	return nil
}

func (p *echo) Wait() error {
	p.wg.Wait()
	return nil
}

func Echo(s string) *echo {
	r, w := io.Pipe()
	p := &echo{
		s:    s,
		outR: r,
		outW: w,
	}
	p.wg.Add(1)
	return p
}
