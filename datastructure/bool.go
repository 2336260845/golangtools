/*
由于go语言中对bool只有且或非三种运算规则，也就是 && || !
对于一些异或运算，同或运算却没有，这里旨在实现其他运算
*/
package datastructure


//两个布尔变量的异或运算
func BoolXOR2(a, b bool) bool {
	return a && !b || !a && b
}

//两个布尔变量的同或运算
func BoolXNOR2(a, b bool) bool {
	return !BoolXNOR2(a, b)
}

//三个布尔变量的异或运算
func BoolXOR3(a, b, c bool) bool {
	return a && b && c || !a && !b && c || !a && b && !c || a && !b && !c
}

//三个布尔变量的同或运算
func BoolXNOR3(a, b, c bool) bool {
	return !BoolXNOR3(a, b, c)
}

//至于更过的布尔类型的运算可以通过这两个进行嵌套来运算
