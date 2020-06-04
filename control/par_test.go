package control

import (
	"errors"
	"testing"

	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/execution"
)

func Test_par(t *testing.T) {
	e := errors.New("e")
	p1 := builtin.Noop()
	p2 := builtin.Failure(e)
	p := Par(p1, p2)
	r := execution.Run(p)
	if r.Err == nil {
		t.Errorf("unexpected success")
	}
}
