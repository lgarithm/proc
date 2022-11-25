/*
make
./bin/test-ssh 2>err.log >out.log
*/
package main

import (
	"os"
	"time"

	"github.com/lgarithm/proc"
	"github.com/lgarithm/proc/builtin"
	"github.com/lgarithm/proc/execution"
	"github.com/lgarithm/proc/iostream"
)

var pwd, _ = os.Getwd()

func main() {
	q := proc.Proc{
		Prog: pwd + "/bin/task",
		Host: `localhost`,
	}
	p := builtin.SSH(q).Timeout(100 * time.Millisecond)
	r := execution.Run(p, &iostream.Std)
	r.Unwrap()
}
