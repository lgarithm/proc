package proc

import (
	"context"

	"github.com/lgarithm/proc/builtin"
	"github.com/lgarithm/proc/control"
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/iostream"
	"github.com/lgarithm/proc/proc"
	"github.com/lgarithm/proc/remote"
)

type (
	Proc = proc.Proc
	P    = execution.P
)

var (
	Main   = execution.Main
	Run    = execution.Run
	Output = execution.Output
	Out    = Output

	Par = control.Par
	Seq = control.Seq

	Echo  = builtin.Echo
	Error = builtin.Failure
	Fn    = builtin.Fn
	FnOk  = builtin.FnOk
	Noop  = builtin.Noop
	Shell = builtin.Shell

	Ignore = control.Ignore
	Lambda = control.Lambda
	Lmd    = Lambda
	Tee    = control.Tee
	Term   = control.Term
	Try    = control.Try
	TryI   = control.TryI
)

var (
	RandomFailure = builtin.RandomFailure
)

var (
	Stdio = iostream.Std
)

type (
	User     = remote.User
	UserHost = remote.UserHost
)

var (
	At   = remote.At
	RPC  = remote.RPC
	SH   = Shell
	SSH  = remote.SSH
	Trpc = remote.Trpc
	Urpc = remote.Urpc
)

func PC(prog string, args ...string) P {
	return Psh(Proc{Prog: prog, Args: args})
}

func Psh(p Proc) P { return SH(p.CmdCtx(context.TODO())) }

func Ps(p ...P) []P { return p }

type Local struct{}

func (Local) PC(prog string, args ...string) P { return PC(prog, args...) }

type CreateP interface {
	PC(prog string, args ...string) P
}
