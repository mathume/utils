package utils

import (
	. "launchpad.net/gocheck"
	"testing"
)

type FUNC struct{}

var _ = Suite(&FUNC{})

func Test(t *testing.T) { TestingT(t) }

func (f *FUNC) TestShortestFirst(c *C) {

	shorterString := "a"
	longerString := "bc"
	expected := longerString

	shorterString, longerString = ShortestFirst(longerString, shorterString)

	c.Assert(longerString, Equals, expected)
}

func (f *FUNC) TestDiffIndex_whenEqualStrings(c *C) {

	anyString := "aa"
	_, err := DiffIndex(anyString, anyString)
	c.Assert(err, Not(IsNil))
}

func (f *FUNC) TestDiffIndex_whenSubstring(c *C) {

	subString := "aa"
	superString := subString + "b"
	expectedIndex := 2

	actualIndex, _ := DiffIndex(subString, superString)

	c.Assert(actualIndex, Equals, expectedIndex)
}

func (f *FUNC) TestDiffIndex_FindsDifference(c *C) {

	s1 := "abcd"
	s2 := "accd"
	expectedIndex := 1

	actualIndex, _ := DiffIndex(s1, s2)

	c.Assert(actualIndex, Equals, expectedIndex)
}

func (f *FUNC) TestIncrement(c *C) {
	i := 0
	i++
	c.Assert(i, Equals, 1)
}

func (f *FUNC) TestContains(c *C) {
	st := []string{"RIPEMD160", "SHA256", "SHA1", "MD5"}
	if Contains(st, "") {
		c.Fatal("Found empty string")
	}
	if !Contains(st, "SHA256") {
		c.Fatal("Didn't find SHA256")
	}
}

func (f *FUNC) TestMax(c *C) {
	max, min := 100, -100
	c.Assert(Max(max, min), Equals, max)
	c.Assert(Max(min, max), Equals, max)
}
