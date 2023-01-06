package proc

import (
	"bytes"
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
)

func Capture(p P) ([]byte, []byte, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	stdpipe := &iostream.StdWriters{
		Stdout: stdout,
		Stderr: stderr,
	}
	r := Run(p, stdpipe)
	return stdout.Bytes(), stderr.Bytes(), r.Err
}

var (
	Par = control.Par
	Seq = control.Seq

	Call = proc.Call

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

type (
	User     = remote.User
	UserHost = remote.UserHost
)

var (
	At     = remote.At
	RPC    = remote.RPC
	SH     = Shell
	SSH    = remote.SSH
	SSHVia = remote.SSHVia
	Trpc   = remote.Trpc
	Urpc   = remote.Urpc
)

func PC(prog string, args ...string) P { return Psh(Call(prog, args...)) }

func Psh(p Proc) P { return SH(p.CmdCtx(context.TODO())) }

func Ps(p ...P) []P { return p }

type Local struct{}

func (Local) PC(prog string, args ...string) P { return PC(prog, args...) }

type LocalDir string

func (d LocalDir) PC(prog string, args ...string) P {
	p := Call(prog, args...)
	p.Dir = string(d)
	return Psh(p)
}

type CreateP interface {
	PC(prog string, args ...string) P
}

type Env = proc.Env

func SetEnv(p *Proc, k, v string) {
	e := make(Env)
	e[k] = v
	p.Env = proc.Merge(p.Env, e)
}

func If(ok bool, p P) P {
	if ok {
		return p
	}
	return Noop()
}

var (
	Stdio = iostream.Std
	Open2 = iostream.Open2
)

type PS1 string

func (ps PS1) Term(p P) P { return Term(string(ps), p) }

type CreatePFn = func(prog string, args ...string) P

func Jump(a UserHost, relay UserHost) CreatePFn {
	return func(prog string, args ...string) P {
		return relay.RPC(a.Call(prog, args...))
	}
}

func WithTerm(pc CreatePFn) CreatePFn {
	return func(prog string, args ...string) P {
		return Term(prog+` % `, pc(prog, args...))
	}
}
