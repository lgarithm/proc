package plugins

import (
	"context"

	"github.com/lgarithm/proc"
	"github.com/lgarithm/proc/builtin"
	"github.com/lgarithm/proc/control"
	"github.com/lgarithm/proc/execution"
)

type (
	Proc = proc.Proc
	P    = execution.P
)

var (
	SH   = builtin.Shell
	SSH  = builtin.SSH
	Term = control.Term
)

type UserHost struct{ User, Host string }

func At(u, h string) UserHost {
	return UserHost{
		User: u,
		Host: h,
	}
}

func PC(prog string, args ...string) P {
	return Psh(Proc{Prog: prog, Args: args})
}

func RPC(host string, prog string, args ...string) P {
	return SSH(Proc{Prog: prog, Args: args, Host: host})
}

func Urpc(uh UserHost, prog string, args ...string) P {
	return SSH(Proc{Prog: prog, Args: args, Host: uh.Host, User: uh.User})
}

func Trpc(ps1, host string, prog string, args ...string) P {
	return Term(ps1, RPC(host, prog, args...))
}

func Psh(p Proc) P { return SH(p.CmdCtx(context.TODO())) }

func Ps(p ...P) []P { return p }
