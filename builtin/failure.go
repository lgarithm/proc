package builtin

import "math/rand"

type failure struct {
	noop
	err error
}

func (p *failure) Wait() error { return p.err }

func Failure(err error) *failure { return &failure{err: err} }

type randomFailure struct {
	noop
	err error
	p   float32
	r   *rand.Rand
}

func (p *randomFailure) Wait() error {
	if p.r.Float32() < p.p {
		return p.err
	}
	return nil
}

func RandomFailure(err error, p float32, r *rand.Rand) *randomFailure {
	return &randomFailure{err: err, p: p, r: r}
}
