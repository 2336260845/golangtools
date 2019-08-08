package datastructure

func MapReverseString2(m map[string]string) (r map[string]string) {
	if m == nil || len(m) == 0 {
		return m
	}
	r = make(map[string]string)
	for i, v := range m {
		r[v] = i
	}
	return r
}
