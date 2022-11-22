package builtin

import (
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"path"
	"time"

	"golang.org/x/crypto/ssh"
)

type sshell struct {
	p       Proc
	timeout time.Duration
	err     error
	client  *ssh.Client
	sess    *ssh.Session

	pty bool
}

func (p *sshell) Stdpipe() (io.Reader, io.Reader, error) {
	if p.client, p.err = newClient(p.p, p.timeout); p.err != nil {
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

func (p *sshell) Timeout(timeout time.Duration) *sshell {
	p.timeout = timeout
	return p
}

// Pty makes session to RequestPty
func (p *sshell) Pty() *sshell {
	p.pty = true
	return p
}

func SSH(p Proc) *sshell { return &sshell{p: p} }

func newClient(p Proc, timeout time.Duration) (*ssh.Client, error) {
	key, err := defaultKeyFile()
	if err != nil {
		return nil, err
	}
	cfg := &ssh.ClientConfig{
		User: getUser(p.User),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}
	client, err := ssh.Dial("tcp", addDefaultPort(p.Host), cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func defaultKeyFile() (ssh.Signer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	file := path.Join(home, ".ssh", "id_rsa")
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ssh.ParsePrivateKey(buf)
}

func addDefaultPort(host string) string {
	_, _, err := net.SplitHostPort(host)
	if err == nil {
		return host
	}
	const defaultPort = "22"
	return net.JoinHostPort(host, defaultPort)
}

func getUser(usr string) string {
	if len(usr) > 0 {
		return usr
	}
	if u, err := user.Current(); err == nil {
		return u.Username
	}
	return ""
}
