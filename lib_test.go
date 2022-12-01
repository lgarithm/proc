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
	isP(proc.Noop())
	isP(proc.Fn(ok))
	isP(proc.FnOk(void))
	isP(proc.Error(nil))

	isCombinator(proc.Par)
	isCombinator(proc.Seq)

	isFEndo(proc.Try)
	isFEndo(proc.Lmd)

	isEndo(proc.Ignore)
}

func isP(P) {}

func isCombinator(func(ps ...P) P) {}

func isFEndo(func(func() P) P) {}

func isEndo(func(P) P) {}

func ok() error { return nil }

func void() {}
