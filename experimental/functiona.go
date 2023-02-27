package experimental

func Dist(f func(string) P, ips []string) []P {
	var ps []P
	for _, ip := range ips {
		ps = append(ps, f(ip))
	}
	return ps
}
