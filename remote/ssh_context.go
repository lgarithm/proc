package remote

import "context"

type sshellContext struct {
	sshell
	ctx context.Context
}

func SSHCtx(p Proc, ctx context.Context) *sshellContext {
	return &sshellContext{
		sshell: sshell{p: p},
		ctx:    ctx,
	}
}

func (p *sshellContext) Start() error {
	go func() {
		<-p.ctx.Done()
		p.sshell.client.Close()
	}()
	return p.sshell.Start()
}
