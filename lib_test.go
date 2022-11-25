package proc_test

import (
	"testing"

	"github.com/lgarithm/proc"
)

type (
	Proc = proc.Proc
	P    = proc.P
)

func Test_1(t *testing.T) {
	isCombinator(proc.Par)
	isCombinator(proc.Seq)

	isFEndo(proc.Try)
	isFEndo(proc.Lmd)

	isEndo(proc.Ignore)
}

func isCombinator(f func(ps ...P) P) {}

func isFEndo(f func(func() P) P) {}

func isEndo(f func(P) P) {}
