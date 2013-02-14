package utils

import (
	"errors"
)

func DiffIndex(s1, s2 string) (diffIndex int, err error) {
	s1, s2 = ShortestFirst(s1, s2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffIndex = i
			return
		}
	}
	if len(s1) == len(s2) {
		err = errors.New("Both strings are equal.")
	} else {
		diffIndex = len(s1)
	}
	return
}

func ShortestFirst(s1, s2 string) (string, string) {
	if len(s1) > len(s2) {
		return s2, s1
	}
	return s1, s2
}

func Contains(list []string, item string) (b bool) {
	for _, v := range list {
		if v == item {
			b = true
		}
	}
	return
}
