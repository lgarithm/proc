package remote

import (
	"io"
	"net"

	"golang.org/x/crypto/ssh"
)

type relayShell struct {
	relay       UserHost // TODO: support multiple relay
	relayClient *ssh.Client
	*sshell
}

func (p *relayShell) Stdpipe() (io.Reader, io.Reader, error) {
	key, err := defaultKeyFile()
	if err != nil {
		return nil, nil, err
	}
	u := p.relay.WithKey(key)
	if p.relayClient, p.err = sshDialFirst(u.Host, u.SSHConfig(p.dialTimeout)); p.err != nil {
		return nil, nil, p.err
	}
	if p.client, p.err = sshDialNext(p.relayClient, p.p.Host, userKey(p.p.User, key, p.dialTimeout)); p.err != nil {
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

func (p *relayShell) Wait() error {
	defer p.relayClient.Close()
	defer p.client.Close()
	return p.sess.Wait()
}

func SSHVia(relay UserHost, p Proc) *relayShell {
	return &relayShell{
		relay:  relay,
		sshell: SSH(p),
	}
}

func sshDialWith(conn net.Conn, host string, cfg *ssh.ClientConfig) (*ssh.Client, error) {
	ncc, chans, reqs, err := ssh.NewClientConn(conn, host, cfg)
	if err != nil {
		return nil, err
	}
	return ssh.NewClient(ncc, chans, reqs), nil
}

func sshDialNext(client *ssh.Client, host string, cfg *ssh.ClientConfig) (*ssh.Client, error) {
	conn, err := client.Dial(`tcp`, addDefaultPort(host))
	if err != nil {
		return nil, err
	}
	return sshDialWith(conn, host, cfg)
}
