package builtin

import (
	"bytes"
	"io"
	"sync"
)

type fn struct {
	f   func() error
	err error
	wg  sync.WaitGroup
}

func (p *fn) Stdpipe() (io.Reader, io.Reader, error) {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	return out, err, nil
}

func (p *fn) Start() error {
	go func() {
		p.err = p.f()
		p.wg.Done()
	}()
	return nil
}

func (p *fn) Wait() error {
	p.wg.Wait()
	return p.err
}

func Fn(f func() error) *fn {
	p := &fn{f: f}
	p.wg.Add(1)
	return p
}

func FnOk(f func()) *fn {
	return Fn(func() error {
		f()
		return nil
	})
}
