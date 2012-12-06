package utils

import (
	"errors"
)

func DiffIndex(s1, s2 string) (diffIndex int, err error) {
	s1, s2 = ShortestFirst(s1, s2)
	for i:=0; i<len(s1); i++ {
		if s1[i] != s2[i] {
			diffIndex = i
			return
		}
	}
	if len(s1) == len(s2){
		err = errors.New("Both strings are equal.")
	}else{
		diffIndex = len(s1)
	}
	return
}

func ShortestFirst(s1, s2 string) (s3, s4 string) {
	if len(s1) > len(s2) {
		s3 = s2
		s4 = s1
	} else {
		s3 = s1
		s4 = s2
	}
	return
}

func HaveEqualContents(s1, s2 []byte) (b bool) {
	n1, n2 := len(s1), len(s2)

	if n1 != n2 {
		return
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return
		}
	}
	b = true
	return
}

func Max(n1, n2 int) (m int){
	m = n1
	if n1 < n2 {
		m = n2
	}
	return	
}

func Contains(list []string, item string) (b bool) {
	for _, v := range list {
		if v == item {
			b = true
		}
	}
	return
}