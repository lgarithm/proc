package main

import (
	"context"
	"log"
	"time"

	"github.com/lgarithm/proc"
	"github.com/lgarithm/proc/remote"
)

func main() {
	f1(1 * time.Second)
	f1(5 * time.Second)
	f1(7 * time.Second)
}

func f1(timeout time.Duration) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	p := proc.Call(`sleep`, `3`)
	p.Host = `127.0.0.1`
	q := remote.SSHCtx(p, ctx)
	r := proc.Run(q, &proc.Stdio)
	if r.Err != nil {
		log.Printf("failed: %v", r.Err)
	} else {
		log.Printf("ok")
	}
}
