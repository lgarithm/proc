package builtin

import (
	"io"
	"os/exec"
)

type shell struct {
	c *exec.Cmd
}

func (p *shell) Stdpipe() (io.Reader, io.Reader, error) {
	stdout, err := p.c.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := p.c.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	return stdout, stderr, nil
}

func (p *shell) Start() error { return p.c.Start() }

func (p *shell) Wait() error { return p.c.Wait() }

func Shell(c *exec.Cmd) *shell { return &shell{c: c} }
