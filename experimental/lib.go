package experimental

import (
	"fmt"
	"time"

	"github.com/lgarithm/proc"
)

type (
	Proc = proc.Proc
	P    = proc.P
	At   = proc.UserHost
)

var (
	lmd  = proc.Lmd
	seq  = proc.Seq
	try  = proc.Try
	echo = proc.Echo
	fnOk = proc.FnOk
	urpc = proc.Urpc
	psh  = proc.Psh
	ps   = proc.Ps
)

func Sleep(d time.Duration) P {
	return fnOk(func() {
		time.Sleep(d)
	})
}

func WaitSSH(a At) P {
	return try(func() P {
		return seq(
			echo(`waiting: `+a.Host),
			Sleep(1*time.Second),
			urpc(a, `pwd`),
		)
	})
}

type CreatePFn = proc.CreatePFn

func WithLog(pc CreatePFn) CreatePFn {
	return func(prog string, args ...string) P {
		var t0 time.Time
		return seq(
			lmd(func() P {
				t0 = time.Now()
				return echo(`BGN ` + prog)
			}),
			pc(prog, args...),
			lmd(func() P {
				d := time.Since(t0)
				return echo(`END ` + prog + fmt.Sprintf(" | took %s", d))
			}),
		)
	}
}
