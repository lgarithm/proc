package execution_test

import (
	"context"
	"testing"

	"github.com/lgarithm/proc-experimental"

	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/execution"
)

func Test_Output(t *testing.T) {
	p := proc.Proc{
		Prog: `echo`,
		Args: []string{`pong`},
	}
	bs := execution.Output(builtin.Shell(p.CmdCtx(context.TODO())))
	want := "pong\n"
	if got := string(bs); got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
