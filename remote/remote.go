package remote

import (
	"github.com/lgarithm/proc/control"
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/proc"
)

type (
	Proc = proc.Proc
	P    = execution.P
)

var (
	term = control.Term
)

type UserHost struct{ User, Host string }

func At(u, h string) UserHost {
	return UserHost{
		User: u,
		Host: h,
	}
}

func (a UserHost) PC(prog string, args ...string) P {
	return Urpc(a, prog, args...)
}

func RPC(host string, prog string, args ...string) P {
	return SSH(Proc{Prog: prog, Args: args, Host: host})
}

func Urpc(a UserHost, prog string, args ...string) P {
	return SSH(Proc{Prog: prog, Args: args, Host: a.Host, User: a.User})
}

func Trpc(ps1, host string, prog string, args ...string) P {
	return term(ps1, RPC(host, prog, args...))
}

type User string

func (u User) At(host string) UserHost { return At(string(u), host) }
