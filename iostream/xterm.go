package iostream

import (
	"fmt"
	"io"
	"os"

	"github.com/lgarithm/proc/xterm"
)

type PromStr string

func (p PromStr) NewTerm(w io.Writer) *Terminal {
	return NewTerminal(string(p), w)
}

type Terminal struct {
	prefix string
	w      io.Writer
}

func (t Terminal) Write(bs []byte) (int, error) {
	fmt.Fprintf(t.w, "%s%s", t.prefix, string(bs))
	return len(bs), nil
}

func NewTerminal(prefix string, w io.Writer) *Terminal {
	return &Terminal{
		prefix: prefix,
		w:      w,
	}
}

func NewTerminalRedirector(prefix string) *StdWriters {
	return &StdWriters{
		Stdout: NewTerminal(prefix, os.Stdout),
		Stderr: NewTerminal(prefix, os.Stderr),
	}
}

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
