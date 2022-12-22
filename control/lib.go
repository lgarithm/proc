package control

import (
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/iostream"
)

type (
	P = execution.P

	StdReaders = iostream.StdReaders
	StdWriters = iostream.StdWriters

	PromStr = iostream.PromStr
	PromoFn = iostream.PromoFn
)

var (
	mix = iostream.Mix
)
