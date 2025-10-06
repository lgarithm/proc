package remote

import (
	"errors"
	"os"
	"os/user"
	"path"

	"golang.org/x/crypto/ssh"
)

func createClientConfig(user string, c config) *ssh.ClientConfig {
	var auth []ssh.AuthMethod
	if len(c.Passwd) > 0 {
		auth = append(auth, ssh.Password(c.Passwd))
	}
	if len(auth) == 0 {
		key, err := defaultKeyFile()
		if err == nil {
			auth = append(auth, ssh.PublicKeys(key))
		}
	}
	return &ssh.ClientConfig{
		User:            getUser(user),
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         c.dialTimeout,
	}
}

func userKey(user string, key ssh.Signer) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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

func getUser(usr string) string {
	if len(usr) > 0 {
		return usr
	}
	if u, err := user.Current(); err == nil {
		return u.Username
	}
	return ""
}
