package utils

import (
    . "launchpad.net/gocheck"
    "testing"
)

var logger = NewFakeWriter()

func Test(t *testing.T) {
	TestingT(t)
}

type FakeConnTests struct{}

var _ = Suite(&FakeConnTests{})

func (s *FakeConnTests) TestFakeConn(c *C) {
	//create
	f := NewFakeConn()
	c.Assert(f.Closed(), Equals, false)
	c.Assert(f.ReturnsErrorOnClose(), Equals, false)
	//send
	msg := []byte("hello")
	n, err := f.Write(msg)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, len(msg))
	act, exp := string(f.Received()), string(msg)
	c.Assert(act, Equals, exp)
	//receive
	bytes := make([]byte, len(msg)+2)
	n, err = f.Read(bytes)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, len(msg))
	c.Assert(string(bytes[:n]), Equals, string(msg))
	//close
	f.ReturnErrorOnClose(true)
	c.Assert(f.ReturnsErrorOnClose(), Equals, true)
	err = f.Close()
	c.Assert(err, NotNil)
	c.Assert(f.Closed(), Equals, true)
}


