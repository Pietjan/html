package html

import (
	"bytes"
	"io"
)

type indentWriter struct {
	w      io.Writer
	prefix []byte

	wroteIndent bool
}

func Tab(repeat int) func(*indentWriter) {
	return Indent('\t', repeat)
}

func Space(repeat int) func(*indentWriter) {
	return Indent(' ', repeat)
}

func Indent(c rune, repeat int) func(*indentWriter) {
	return func(w *indentWriter) {
		w.prefix = bytes.Repeat([]byte{byte(c)}, repeat)
	}
}

func IndentWriter(w io.Writer, options ...func(*indentWriter)) io.Writer {
	iw := indentWriter{
		w:      w,
		prefix: bytes.Repeat([]byte{' '}, 2),
	}

	for _, option := range options {
		option(&iw)
	}

	return &iw
}

func (iw *indentWriter) Write(p []byte) (n int, err error) {
	for i, b := range p {
		if b == '\n' {
			iw.wroteIndent = false
		} else {
			if !iw.wroteIndent {
				_, err = iw.w.Write(iw.prefix)
				if err != nil {
					return n, err
				}
				iw.wroteIndent = true
			}
		}
		_, err = iw.w.Write(p[i : i+1])
		if err != nil {
			return n, err
		}
		n++
	}
	return len(p), nil
}
