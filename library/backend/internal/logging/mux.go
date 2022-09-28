package logging

import "io"

var _ io.Writer = &writerMux{}

type writerMux struct {
	writers []io.Writer
}

func (me *writerMux) Write(p []byte) (n int, err error) {
	lastN := 0
	for _, w := range me.writers {
		n, err := w.Write(p)
		if err != nil {
			return n, err
		}
		lastN = n
	}
	return lastN, nil
}
