package stringutil

import (
	"bytes"
	"strings"
)

type StringWriter struct {
	b bytes.Buffer
}

func NewStringWriter() *StringWriter {
	return &StringWriter{}
}

func (w *StringWriter) Write(p []byte) (n int, err error) {
	return w.b.Write(p)
}

func (w *StringWriter) String() string {
	return strings.TrimSpace(w.b.String())
}
