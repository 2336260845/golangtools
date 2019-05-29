package datastructure

import "testing"

func TestBoolXOR2(t *testing.T) {
	funcName := "TestBoolXOR2"
	if BoolXOR2(true, true) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR2(false, false) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR2(true, false) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR2(false, true) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
}

func TestBoolXNOR2(t *testing.T) {
	//如果两个布尔变量的异或运算没问题，那么同或运算也没问题，因为同或运算只是在异或运算上取反
}

func TestBoolXOR3(t *testing.T) {
	funcName := "TestBoolXOR3"
	if BoolXOR3(true, true, true) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(true, true, false) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(true, false, true) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(false, true, true) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(true, false, false) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(false, false, true) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(false, true, false) != true {
		t.Errorf("Test %s error", funcName)
		return
	}
	if BoolXOR3(false, false, false) != false {
		t.Errorf("Test %s error", funcName)
		return
	}
}

func TestBoolXNOR3(t *testing.T) {
	//如果三个布尔变量的异或运算没问题，那么同或运算也没问题，因为同或运算只是在异或运算上取反
}
