package utils

import (
    . "launchpad.net/gocheck"
)

type WRITER struct{}

var _ = Suite(&WRITER{})

func (w *WRITER)TestWrite(c *C){
	fw := NewFakeWriter()
	msg := []byte("my message")
	n, err := fw.Write(msg)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, len(msg))
	c.Assert(string(msg), Equals, fw.Msg())
	fw.Write(msg)
	c.Assert(string(msg) + string(msg), Equals, fw.Msg())
}

func (w *WRITER)TestSetMsg(c *C){
	fw := NewFakeWriter()
	fw.Write([]byte("any message of non-zero length"))
	fw.SetMsg("")
	c.Assert(fw.Msg(), Equals, "")
}
