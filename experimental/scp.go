package experimental

func Scp(uh At, local, remote string) P {
	p0 := Proc{
		Prog: `scp`,
		Args: []string{
			`-o`, `StrictHostKeyChecking=no`,
			local,
			uh.User + `@` + uh.Host + ":" + remote,
		},
	}
	return seq(
		psh(p0),
		echo("done scp: "+local+" to "+uh.Host+":"+remote),
	)
}
