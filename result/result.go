package result

import (
	"fmt"
	"os"
	"time"
)

type Result struct {
	Err  error
	Took time.Duration
}

func (r Result) Unwrap() {
	if r.Err != nil {
		fmt.Fprintf(os.Stderr, "failed %v\n", r.Err)
		os.Exit(1)
		return
	}
	fmt.Printf("took %s\n", r.Took)
}

func Return(t0 time.Time, err error) Result {
	return Result{
		Err:  err,
		Took: time.Since(t0),
	}
}
