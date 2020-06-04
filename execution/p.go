package execution

import (
	"io"
	"time"

	"github.com/lgarithm/proc-experimental/iostream"
	"github.com/lgarithm/proc-experimental/result"
)

type P interface {
	Stdpipe() (io.Reader, io.Reader, error)
	Start() error
	Wait() error
}

func Run(p P, redirectors ...*iostream.StdWriters) result.Result {
	t0 := time.Now()
	stdout, stderr, err := p.Stdpipe()
	if err != nil {
		return result.Return(t0, err)
	}
	results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
	ioDone := results.Stream(redirectors...)
	if err := p.Start(); err != nil {
		return result.Return(t0, err)
	}
	ioDone.Wait()
	return result.Return(t0, p.Wait())
}
