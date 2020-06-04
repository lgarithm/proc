package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/control"
	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
	"github.com/lgarithm/proc-experimental/xterm"
)

func main() {
	flag.Parse()
	parExample()
	seqExample()
	tryExample()
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
		if r := execution.Run(p, w); r.Err != nil {
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
		if r := execution.Run(p, w); r.Err != nil {
			fmt.Printf("failed: %v\n", r.Err)
		}
	}
}

func tryExample() {
	e := errors.New("e")
	var n int
	q := func() execution.P {
		n++
		fmt.Printf("trial #%d\n", n)
		return control.Par(
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
			builtin.RandomFailure(e, 0.9, rand.New(rand.NewSource(time.Now().UnixNano()))),
		)
	}
	p := control.Try(q)
	w := iostream.NewXTermRedirector(`x`, xterm.Green)
	if r := execution.Run(p, w); r.Err != nil {
		fmt.Printf("failed: %v\n", r.Err)
	}
}
