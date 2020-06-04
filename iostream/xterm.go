package iostream

import (
	"fmt"
	"io"
	"os"

	"github.com/lgarithm/proc-experimental/xterm"
)

type XtermWriter struct {
	prefix string
	w      io.Writer
}

func (x XtermWriter) Write(bs []byte) (int, error) {
	fmt.Fprintf(x.w, "[%s] %s", x.prefix, string(bs))
	return len(bs), nil
}

func NewXTermRedirector(name string, c xterm.Color) *StdWriters {
	if c == nil {
		c = xterm.NoColor
	}
	return &StdWriters{
		Stdout: &XtermWriter{
			prefix: c.S(name) + "::stdout",
			w:      os.Stdout,
		},
		Stderr: &XtermWriter{
			prefix: c.S(name) + "::" + xterm.Warn.S("stderr"),
			w:      os.Stderr,
		},
	}
}
