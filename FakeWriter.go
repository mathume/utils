package utils

import (
	"io"
)

type FakeWriter interface{
	io.Writer
	Msg() string
	SetMsg(msg string)
}

type Writer struct {
	msg string
}

func (w *Writer) Clear() {
	w.msg = ""
	return
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.msg += string(p)
	n, err = len(p), nil
	return
}

func (w *Writer)Msg() string{
	return w.msg
}

func (w *Writer)SetMsg(msg string){
	w.msg = msg
	return
}

func NewFakeWriter() FakeWriter{
	return new(Writer)
}