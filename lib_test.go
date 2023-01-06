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
	var isP = is[P]
	isP(proc.Noop())
	isP(proc.Fn(ok))
	isP(proc.FnOk(void))
	isP(proc.Error(nil))

	isCombinator(proc.Par)
	isCombinator(proc.Seq)

	isFEndo(proc.Try)
	isFEndo(proc.Lmd)

	var isEndo = isEndoT[P]
	isEndo(proc.Ignore)

	var ps proc.PS1
	isEndo(ps.Term)
}

func is[T any](T) {}

func isCombinator(func(ps ...P) P) {}

func isFEndo(func(func() P) P) {}

func ok() error { return nil }

func void() {}

func Test_2(t *testing.T) {
	var l proc.Local
	isCreateP(l)
	var a proc.UserHost
	isCreateP(a)
	var d proc.LocalDir
	isCreateP(d)
}

func isCreateP(proc.CreateP) {}

func Test_3(t *testing.T) {
	var isEndo = isEndoT[proc.CreatePFn]
	isEndo(proc.WithTerm)
}

func isEndoT[T any](func(T) T) {}
