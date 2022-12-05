package proc

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Proc represents a general purpose process
type Proc struct {
	Name string
	Prog string
	Args []string
	Env  Env
	Dir  string
	Host string
	User string
}

func (p Proc) CmdCtx(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, p.Prog, p.Args...)
	cmd.Env = updatedEnvFrom(p.Env, os.Environ())
	cmd.Dir = p.Dir
	return cmd
}

func (p Proc) Script() string {
	buf := &bytes.Buffer{}
	var chdir string
	if len(p.Dir) > 0 {
		chdir = fmt.Sprintf("-C %s", p.Dir)
	}
	fmt.Fprintf(buf, "env %s\\\n", chdir)
	for k, v := range p.Env {
		fmt.Fprintf(buf, "\t%s=%q \\\n", k, v)
	}
	fmt.Fprintf(buf, "\t%s \\\n", p.Prog)
	for _, a := range p.Args {
		fmt.Fprintf(buf, "\t%s \\\n", a)
	}
	fmt.Fprintf(buf, "\n")
	return buf.String()
}

func Call(prog string, args ...string) Proc { return Proc{Prog: prog, Args: args} }
