package builtin

type failure struct {
	noop
	err error
}

func (p *failure) Wait() error { return p.err }

func Failure(err error) *failure { return &failure{err: err} }
