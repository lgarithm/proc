package control

import (
	"bytes"
	"testing"

	"github.com/lgarithm/proc-experimental/builtin"
	"github.com/lgarithm/proc-experimental/execution"
	"github.com/lgarithm/proc-experimental/iostream"
)

func Test_term(t *testing.T) {
	p1 := Term(`[1] `, builtin.Echo(`pong`))
	out := &bytes.Buffer{}
	stdpipe := &iostream.StdWriters{
		Stdout: out,
	}
	r := execution.Run(p1, stdpipe)
	if r.Err != nil {
		t.Errorf("unexpected failure")
	}
	want := "[1] pong\n"
	if got := out.String(); got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
