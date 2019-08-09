package datastructure


//************************************************map操作********************************************************//
//map翻转
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


func MapReverseStringInt(m map[string]int) (r map[int]string) {
	if m == nil {
		return nil
	}
	if len(m) == 0 {
		return map[int]string{}
	}
	r = make(map[int]string)
	for i, v := range m {
		r[v] = i
	}
	return r
}

func MapReverseIntString(m map[int]string) (r map[string]int) {
	if m == nil {
		return nil
	}
	if len(m) == 0 {
		return map[string]int{}
	}
	r = make(map[string]int)
	for i, v := range m {
		r[v] = i
	}
	return r
}
