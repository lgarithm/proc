package iostream

import (
	"fmt"
	"io"
)

type PromoFn func(int) string

func (p PromStr) LineNo() PromoFn {
	return func(i int) string {
		return fmt.Sprintf("%s%d ", p, i)
	}
}

func (p PromoFn) NewTerm(w io.Writer) *DynTerm {
	return NewDynTerm(p, w)
}

type DynTerm struct {
	line int
	ps   PromoFn
	w    io.Writer
}

func (t *DynTerm) Write(bs []byte) (int, error) {
	prefix := t.ps(t.line)
	t.line++
	fmt.Fprintf(t.w, "%s%s", prefix, string(bs))
	return len(bs), nil
}

func NewDynTerm(ps PromoFn, w io.Writer) *DynTerm {
	return &DynTerm{
		ps: ps,
		w:  w,
	}
}
