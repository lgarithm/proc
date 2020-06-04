package local

import (
	"context"
	"io"
	"os/exec"
	"time"

	proc "github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/iostream"
	"github.com/lgarithm/proc-experimental/xterm"
)

func RunWith(p proc.Proc, redirectors ...*iostream.StdWriters) proc.Result {
	t0 := time.Now()
	cmd := p.CmdCtx(context.TODO())
	stdout, stderr, err := stdpipe(cmd)
	if err != nil {
		return proc.Return(t0, err)
	}
	results := iostream.StdReaders{Stdout: stdout, Stderr: stderr}
	ioDone := results.Stream(redirectors...)
	if err := cmd.Start(); err != nil {
		return proc.Return(t0, err)
	}
	ioDone.Wait() // call this before cmd.Wait!
	return proc.Return(t0, cmd.Wait())
}

func DefaultRedirectors(p proc.Proc) []*iostream.StdWriters {
	var redirectors []*iostream.StdWriters
	redirectors = append(redirectors, iostream.NewXTermRedirector(p.Name, xterm.Green))
	return redirectors
}

func stdpipe(cmd *exec.Cmd) (io.Reader, io.Reader, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	return stdout, stderr, nil
}
