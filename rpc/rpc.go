package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"

	"github.com/lgarithm/proc"
)

type (
	P    = proc.P
	Proc = proc.Proc
)

func exe(p Proc) ([]byte, []byte, error) { return proc.Capture(proc.Psh(p)) }

type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type RPC struct {
	Host string
	Port int
}

func (s RPC) Run() error {
	srv := rpc.NewServer()
	if err := srv.Register(s); err != nil {
		return err
	}
	mux := &http.ServeMux{}
	mux.Handle(rpc.DefaultRPCPath, srv)
	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
	hs := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return hs.ListenAndServe()
}

func (RPC) Call(req *Proc, resp *Result) error {
	log.Printf("RPC.Call(%#v)", req)
	o, e, err := exe(*req)
	if err != nil {
		return err
	}
	*resp = Result{
		Stdout: string(o),
		Stderr: string(e),
	}
	return nil
}

func Call(s RPC, args ...string) error {
	if len(args) <= 0 {
		return fmt.Errorf("missing prog")
	}
	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
	cli, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return err
	}
	req := Proc{
		Prog: args[0],
		Args: args[1:],
	}
	var resp Result
	if err := cli.Call("RPC.Call", &req, &resp); err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%s\n", resp.Stdout)
	fmt.Fprintf(os.Stderr, "%s\n", resp.Stderr)
	return nil
}
