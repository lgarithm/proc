package remote

import (
	"errors"
	"io"
	"net"
	"os"
	"os/user"
	"path"
	"time"

	"golang.org/x/crypto/ssh"
)

type sshell struct {
	p           Proc
	dialTimeout time.Duration
	err         error
	client      *ssh.Client
	sess        *ssh.Session

	pty bool
}

func (p *sshell) Stdpipe() (io.Reader, io.Reader, error) {
	if p.client, p.err = newClient(p.p, p.dialTimeout); p.err != nil {
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

func newClient(p Proc, timeout time.Duration) (*ssh.Client, error) {
	key, err := defaultKeyFile()
	if err != nil {
		return nil, err
	}
	client, err := sshDialFirst(p.Host, userKey(getUser(p.User), key, timeout))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func userKey(user string, key ssh.Signer, timeout time.Duration) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}
}

var identityFiles = []string{
	`id_rsa`,
	`id_ecdsa`,
	`id_ecdsa_sk`,
	`id_ed25519`,
	`id_ed25519_sk`,
	`id_dsa`,
}

var errIdentityFileNotFound = errors.New(`identity files not found`)

func defaultKeyFile() (ssh.Signer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	for _, id := range identityFiles {
		file := path.Join(home, ".ssh", id)
		if buf, err := os.ReadFile(file); err == nil {
			return ssh.ParsePrivateKey(buf)
		}
	}
	return nil, errIdentityFileNotFound
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

func sshDialFirst(host string, cfg *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial(`tcp`, addDefaultPort(host), cfg)
}
