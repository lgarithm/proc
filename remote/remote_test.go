package remote

import "testing"

func Test_1(t *testing.T) {
	var s sshell
	isP(&s)

	var r relayShell
	isP(&r)
}

func isP(P) {}
