package remote

import (
	"time"

	"github.com/lgarithm/proc/control"
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/proc"
	"golang.org/x/crypto/ssh"
)

type (
	Proc = proc.Proc
	P    = execution.P
)

var (
	term = control.Term
)

type User string

func (u User) At(host string) UserHost { return At(string(u), host) }

type UserHost struct{ User, Host string }

func At(u, h string) UserHost {
	return UserHost{
		User: u,
		Host: h,
	}
}

func (a UserHost) Call(prog string, args ...string) Proc {
	return Proc{Prog: prog, Args: args, Host: a.Host, User: a.User}
}

func (a UserHost) PC(prog string, args ...string) P {
	return Urpc(a, prog, args...)
}

func (a UserHost) RPC(p Proc) P { return SSHVia(a, p) }

func RPC(host string, prog string, args ...string) P {
	return SSH(Proc{Prog: prog, Args: args, Host: host})
}

func Urpc(a UserHost, prog string, args ...string) P { return SSH(a.Call(prog, args...)) }

func Trpc(ps1, host string, prog string, args ...string) P {
	return term(ps1, RPC(host, prog, args...))
}

func (a UserHost) WithKey(key ssh.Signer) UserHostAuth {
	return UserHostAuth{
		User: a.User,
		Host: a.Host,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
}

type UserHostAuth struct {
	User string
	Host string
	Auth []ssh.AuthMethod
}

func (u UserHostAuth) SSHConfig(timeout time.Duration) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:            u.User,
		Auth:            u.Auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}
}
