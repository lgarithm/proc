package proc

import (
	"github.com/lgarithm/proc/builtin"
	"github.com/lgarithm/proc/control"
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/iostream"
	"github.com/lgarithm/proc/plugins"
	"github.com/lgarithm/proc/proc"
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
	Fn    = builtin.Fn
	FnOk  = builtin.FnOk
	Shell = builtin.Shell
	SH    = Shell
	SSH   = builtin.SSH

	Ignore = control.Ignore
	Lambda = control.Lambda
	Lmd    = Lambda
	Tee    = control.Tee
	Term   = control.Term
	Try    = control.Try
	TryI   = control.TryI
)

var (
	Stdio = iostream.Std
)

type UserHost = plugins.UserHost

var (
	At   = plugins.At
	PC   = plugins.PC
	RPC  = plugins.RPC
	Urpc = plugins.Urpc
	Trpc = plugins.Trpc
	Ps   = plugins.Ps
	Psh  = plugins.Psh
)
