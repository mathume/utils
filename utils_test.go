package utils

import (
	"testing"
)

func TestShortestFirst(t *testing.T) {

	shorterString := "a"
	longerString := "bc"
	expected := longerString

	shorterString, longerString = ShortestFirst(longerString, shorterString)

	if longerString != expected {
		t.Fail()
	}
}

func TestDiffIndex_whenEqualStrings(t *testing.T) {

	anyString := "aa"
	_, err := DiffIndex(anyString, anyString)
	if err == nil {
		t.Fail()
	}
}

func TestDiffIndex_whenSubstring(t *testing.T) {

	subString := "aa"
	superString := subString + "b"
	expectedIndex := 2

	actualIndex, _ := DiffIndex(subString, superString)

	if actualIndex != expectedIndex {
		t.Fatal(actualIndex)
	}
}

func TestDiffIndex_FindsDifference(t *testing.T) {

	s1 := "abcd"
	s2 := "accd"
	expectedIndex := 1

	actualIndex, _ := DiffIndex(s1, s2)

	if actualIndex != expectedIndex {
		t.Fatal(actualIndex)
	}
}

func TestIncrement(t *testing.T) {
	i := 0
	i++
	if i != 1 {
		t.Fail()
	}
}
