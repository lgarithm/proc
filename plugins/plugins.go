package plugins

import (
	"context"

	"github.com/lgarithm/proc-experimental"
	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/control"
	"github.com/lgarithm/proc-experimental/execution"
)

type (
	Proc = proc.Proc
	P    = execution.P
)

var (
	Run = execution.Run
	Out = execution.Output

	Par = control.Par
	Seq = control.Seq

	Echo = builtin.Echo
	Fn   = builtin.Fn
	FnOk = builtin.FnOk
	SH   = builtin.Shell
	SSH  = builtin.SSH

	Ignore = control.Ignore
	Lmd    = control.Lambda
	Tee    = control.Tee
	Term   = control.Term
	Try    = control.Try
	TryI   = control.TryI
)

type UserHost struct{ User, Host string }

func At(u, h string) UserHost {
	return UserHost{
		User: u,
		Host: h,
	}
}

func PC(prog string, args ...string) P {
	return Psh(proc.Proc{Prog: prog, Args: args})
}

func RPC(host string, prog string, args ...string) P {
	return SSH(proc.Proc{Prog: prog, Args: args, Host: host})
}

func Urpc(uh UserHost, prog string, args ...string) P {
	return SSH(proc.Proc{Prog: prog, Args: args, Host: uh.Host, User: uh.User})
}

func Trpc(ps1, host string, prog string, args ...string) P {
	return Term(ps1, RPC(host, prog, args...))
}

func Psh(p proc.Proc) P { return SH(p.CmdCtx(context.TODO())) }

func Ps(p ...P) []P { return p }
