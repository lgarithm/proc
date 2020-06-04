package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/control"
	"github.com/lgarithm/proc-experimental/iostream"
	"github.com/lgarithm/proc-experimental/xterm"
)

func main() {
	flag.Parse()
	parExample()
	seqExample()
}

func parExample() {
	p := proc.Proc{
		Prog: `task`,
		Args: []string{`par-example`},
	}
	{
		p := control.Par(
			builtin.Shell(p.CmdCtx(context.TODO())),
			builtin.Shell(p.CmdCtx(context.TODO())),
		)
		w := iostream.NewXTermRedirector(`x`, xterm.Green)
		if r := control.Run(p, w); r.Err != nil {
			fmt.Printf("failed: %v\n", r.Err)
		}
	}
}

func seqExample() {
	p := proc.Proc{
		Prog: `task`,
		Args: []string{`seq-example`},
	}
	{
		p := control.Seq(
			builtin.Shell(p.CmdCtx(context.TODO())),
			builtin.Shell(p.CmdCtx(context.TODO())),
		)
		w := iostream.NewXTermRedirector(`x`, xterm.Green)
		if r := control.Run(p, w); r.Err != nil {
			fmt.Printf("failed: %v\n", r.Err)
		}
	}
}
