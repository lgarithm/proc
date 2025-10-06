package remote

import (
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type config struct {
	dialTimeout time.Duration
	pty         bool
	Passwd      string
}

type sshell struct {
	config
	p      Proc
	err    error
	client *ssh.Client
	sess   *ssh.Session
}

func (p *sshell) Stdpipe() (io.Reader, io.Reader, error) {
	if p.client, p.err = newClient(p.p, p.config); p.err != nil {
		return nil, nil, p.err
	}
	if p.sess, p.err = p.client.NewSession(); p.err != nil {
		return nil, nil, p.err
	}
	var stdout, stderr io.Reader
	if stdout, p.err = p.sess.StdoutPipe(); p.err != nil {
		return nil, nil, p.err
	}
	if stderr, p.err = p.sess.StderrPipe(); p.err != nil {
		return nil, nil, p.err
	}
	// would cause stderr merged into stdout
	if p.pty {
		if p.err = p.sess.RequestPty("xterm", 80, 40, nil); p.err != nil {
			return nil, nil, p.err
		}
	}
	return stdout, stderr, nil
}

func (p *sshell) Start() error { return p.sess.Start(p.p.Script()) }

func (p *sshell) Wait() error {
	defer p.client.Close()
	return p.sess.Wait()
}

func (p *sshell) DialTimeout(timeout time.Duration) *sshell {
	p.dialTimeout = timeout
	return p
}

// Pty makes session to RequestPty
func (p *sshell) Pty() *sshell {
	p.pty = true
	return p
}

func SSH(p Proc) *sshell { return &sshell{p: p} }

func newClient(p Proc, c config) (*ssh.Client, error) {
	cc := createClientConfig(p.User, c)
	client, err := sshDialFirst(p.Host, cc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func addDefaultPort(host string) string {
	_, _, err := net.SplitHostPort(host)
	if err == nil {
		return host
	}
	const defaultPort = "22"
	return net.JoinHostPort(host, defaultPort)
}

func sshDialFirst(host string, cfg *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial(`tcp`, addDefaultPort(host), cfg)
}
