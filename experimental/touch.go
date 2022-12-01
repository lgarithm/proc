package experimental

import (
	"fmt"
	"os"
	"path"
)

func Touch(a At, filepath string, content []byte) P {
	filename := path.Base(filepath)
	return seq(
		lmd(func() P {
			f, err := os.CreateTemp(os.TempDir(), "t-")
			if err != nil {
				panic(err)
			}
			if _, err := f.Write(content); err != nil {
				panic(err)
			}
			ps := ps(Scp(a, f.Name(), filename))
			if filename != filepath {
				ps = append(ps, urpc(a, `sudo`, `mv`, filename, filepath))
			}
			ps = append(ps, echo(fmt.Sprintf(`done mv sudo %s %s`, filename, filepath)))
			return seq(ps...)
		}),
	)
}

func TouchExe(a At, filepath string, content []byte) P {
	return seq(
		Touch(a, filepath, content),
		urpc(a, `sudo`, `chmod`, `+x`, filepath),
	)
}
