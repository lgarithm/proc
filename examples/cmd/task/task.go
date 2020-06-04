package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Fprintf(os.Stdout, "stdout: pid: %d\n", pid)
	fmt.Fprintf(os.Stderr, "stderr: pid: %d\n", pid)
	filename := fmt.Sprintf("%d.log", pid)
	ioutil.WriteFile(filename, []byte(`xx`), os.ModePerm)
}
