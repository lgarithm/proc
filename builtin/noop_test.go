package builtin

import (
	"bytes"
	"testing"

	"github.com/lgarithm/proc-experimental/control"
	"github.com/lgarithm/proc-experimental/iostream"
)

func Test_noop(t *testing.T) {
	p := Noop()
	out := &bytes.Buffer{}
	stdpipe := &iostream.StdWriters{
		Stdout: out,
	}
	control.Run(p, stdpipe)
	if got := out.String(); got != "" {
		t.Errorf("want %q, got %q", "", got)
	}
}
