package utils

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
