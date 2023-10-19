package experimental

func RunScript(a At, script string, filename string, su bool) P {
	run := urpc(a, `./`+filename)
	if su {
		run = urpc(a, `sudo `, `./`+filename)
	}
	return seq(
		TouchExe(a, filename, []byte(script)),
		run,
	)
}
