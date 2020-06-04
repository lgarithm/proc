package execution

import (
	"io"
	"time"

	proc "github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/iostream"
)

type P interface {
	Stdpipe() (io.Reader, io.Reader, error)
	Start() error
	Wait() error
}

func Run(p P, redirectors ...*iostream.StdWriters) proc.Result {
	t0 := time.Now()
	stdout, stderr, err := p.Stdpipe()
	if err != nil {
		return proc.Return(t0, err)
	}
	results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
	ioDone := results.Stream(redirectors...)
	if err := p.Start(); err != nil {
		return proc.Return(t0, err)
	}
	ioDone.Wait()
	return proc.Return(t0, p.Wait())
}
