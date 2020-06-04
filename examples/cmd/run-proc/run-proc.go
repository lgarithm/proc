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

	p := proc.Proc{
		Prog: `task`,
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
