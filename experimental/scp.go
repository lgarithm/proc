package experimental

func Scp(uh At, local, remote string) P {
	p0 := Proc{
		Prog: `scp`,
		Args: []string{local, uh.User + `@` + uh.Host + ":" + remote},
	}
	return seq(
		psh(p0),
		echo("done scp: "+local+" to "+uh.Host+":"+remote),
	)
}
