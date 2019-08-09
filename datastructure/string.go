package datastructure

//************************************************string操作********************************************************//
//判断string中是否包含某几个字符

func StringContainRuneOne(s string, r rune) (b bool, index int) {
	if len(s) <= 0 {
		return false, -1
	}
	for i, ch := range s {
		if rune(ch) == r {
			return true, i
		}
	}
	return false, -1
}

