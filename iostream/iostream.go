package iostream

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

// Tee redirects r to ws
func Tee(r io.Reader, ws ...io.Writer) error {
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		for _, w := range ws {
			fmt.Fprintln(w, string(line))
		}
	}
}

func Mix(w io.Writer, rs ...io.Reader) error {
	var mu sync.Mutex
	writeline := func(line []byte) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintln(w, string(line))
	}
	errs := make([]error, len(rs))
	var wg sync.WaitGroup
	for i, r := range rs {
		wg.Add(1)
		go func(i int, r io.Reader) {
			br := bufio.NewReader(r)
			for {
				line, _, err := br.ReadLine()
				if err != nil {
					if err != io.EOF {
						errs[i] = err
					}
					break
				}
				writeline(line)
			}
			wg.Done()
		}(i, r)
	}
	wg.Wait()
	// FIXME: merge errs
	return nil
}
