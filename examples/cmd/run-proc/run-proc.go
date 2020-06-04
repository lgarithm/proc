package main

import (
	"flag"

	"github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/runner/local"
)

func main() {
	flag.Parse()

	p := proc.Proc{
		Prog: `task`,
	}
	local.RunWith(p, local.DefaultRedirectors(p)...)
}
