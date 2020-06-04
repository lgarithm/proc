package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	cmdline := strings.Join(os.Args, " ")
	pid := os.Getpid()
	fmt.Fprintf(os.Stdout, "stdout: pid: %d $ %s\n", pid, cmdline)
	fmt.Fprintf(os.Stderr, "stderr: pid: %d $ %s\n", pid, cmdline)
	filename := fmt.Sprintf("%d.log", pid)
	ioutil.WriteFile(filename, []byte(`xx`), os.ModePerm)
}
