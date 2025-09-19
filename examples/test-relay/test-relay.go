package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/lgarithm/proc"
)

var (
	h = flag.String(`h`, ``, ``)
	j = flag.String(`j`, ``, ``)
)

func main() {
	flag.Parse()
	d, e := PingJump(*h, *j)
	log.Printf("%s %v", d, e)
}

func PingJump(h, j string) (time.Duration, error) {
	u := os.Getenv(`USER`)
	a := proc.At(u, j)
	q := proc.Proc{
		Prog: `pwd`,
		Host: h,
		User: u,
	}
	p := proc.SSHVia(a, q)
	t0 := time.Now()
	r := proc.Run(p)
	d := time.Since(t0)
	return d, r.Err
}
