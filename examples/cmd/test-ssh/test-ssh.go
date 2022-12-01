/*
make
./bin/test-ssh 2>err.log >out.log
*/
package main

import (
	"os"
	"time"

	"github.com/lgarithm/proc"
	"github.com/lgarithm/proc/iostream"
)

var pwd, _ = os.Getwd()

func main() {
	q := proc.Proc{
		Prog: pwd + "/bin/task",
		Host: `localhost`,
	}
	p := proc.SSH(q).Timeout(100 * time.Millisecond)
	r := proc.Run(p, &iostream.Std)
	r.Unwrap()
}
