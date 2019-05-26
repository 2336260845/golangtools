package datastructure

import (
	"testing"
)

func TestSliceRemoveDupString(t *testing.T) {
	funcName := "TestSliceRemoveDupString"
	s := []string{"mark", "marry", "jack", "jack", "andy", "mark"}
	s, err := SliceRemoveDupString(s)
	if err != nil {
		t.Errorf("Test %s error", funcName)
		return
	}
	if len(s) != 4 {
		t.Errorf("%s error, return slice length is worry", funcName)
		return
	}
	f1 := false
	f2 := false
	f3 := false
	f4 := false
	for _, v := range s {
		if v == "mark" {
			f1 = true
		} else if v == "marry" {
			f2 = true
		} else if v == "jack" {
			f3 = true
		} else if v == "andy" {
			f4 = true
		} else {
			t.Errorf("%s error, exist wrong elenment[%v] in slice", funcName, v)
		}
	}
	if !(f1 && f2 && f3 && f4) {
		t.Errorf("%s error, element is eliminated", funcName)
	}
}

func TestSliceRemoveDupInt(t *testing.T) {
	funcName := "TestSliceRemoveDupInt"
	s := []int{1,2,3,1,2,3,4}
	s, err := SliceRemoveDupInt(s)
	if err != nil {
		t.Errorf("Test %s error", funcName)
		return
	}
	if len(s) != 4 {
		t.Errorf("%s error, return slice length is worry", funcName)
		return
	}
	f1 := false
	f2 := false
	f3 := false
	f4 := false
	for _, v := range s {
		if v == 1 {
			f1 = true
		} else if v == 2 {
			f2 = true
		} else if v == 3 {
			f3 = true
		} else if v == 4 {
			f4 = true
		} else {
			t.Errorf("%s error, exist wrong elenment[%v] in slice", funcName, v)
		}
	}
	if !(f1 && f2 && f3 && f4) {
		t.Errorf("%s error, element is eliminated", funcName)
	}
}

func TestSliceRemoveDupInt64(t *testing.T) {
	funcName := "TestSliceRemoveDupInt64"
	s := []int64{1,2,3,1,2,3,4}
	s, err := SliceRemoveDupInt64(s)
	if err != nil {
		t.Errorf("Test %s error", funcName)
		return
	}
	if len(s) != 4 {
		t.Errorf("%s error, return slice length is worry", funcName)
		return
	}
	f1 := false
	f2 := false
	f3 := false
	f4 := false
	for _, v := range s {
		if v == 1 {
			f1 = true
		} else if v == 2 {
			f2 = true
		} else if v == 3 {
			f3 = true
		} else if v == 4 {
			f4 = true
		} else {
			t.Errorf("%s error, exist wrong elenment[%v] in slice", funcName, v)
		}
	}
	if !(f1 && f2 && f3 && f4) {
		t.Errorf("%s error, element is eliminated", funcName)
	}
}

func TestSliceRemoveDupInt32(t *testing.T) {
	funcName := "TestSliceRemoveDupInt32"
	s := []int32{1,2,3,1,2,3,4}
	s, err := SliceRemoveDupInt32(s)
	if err != nil {
		t.Errorf("Test %s error", funcName)
		return
	}
	if len(s) != 4 {
		t.Errorf("%s error, return slice length is worry", funcName)
		return
	}
	f1 := false
	f2 := false
	f3 := false
	f4 := false
	for _, v := range s {
		if v == 1 {
			f1 = true
		} else if v == 2 {
			f2 = true
		} else if v == 3 {
			f3 = true
		} else if v == 4 {
			f4 = true
		} else {
			t.Errorf("%s error, exist wrong elenment[%v] in slice", funcName, v)
		}
	}
	if !(f1 && f2 && f3 && f4) {
		t.Errorf("%s error, element is eliminated", funcName)
	}
}

func TestSliceRemoveDupByte(t *testing.T) {
	funcName := "TestSliceRemoveDupByte"
	s := []byte{'a', 'b', 'c', 'd', 'a', 'b', 'c'}
	s, err := SliceRemoveDupByte(s)
	if err != nil {
		t.Errorf("Test %s error", funcName)
		return
	}
	if len(s) != 4 {
		t.Errorf("%s error, return slice length is worry", funcName)
		return
	}
	f1 := false
	f2 := false
	f3 := false
	f4 := false
	for _, v := range s {
		if v == 'a' {
			f1 = true
		} else if v == 'b' {
			f2 = true
		} else if v == 'c' {
			f3 = true
		} else if v == 'd' {
			f4 = true
		} else {
			t.Errorf("%s error, exist wrong elenment[%v] in slice", funcName, v)
		}
	}
	if !(f1 && f2 && f3 && f4) {
		t.Errorf("%s error, element is eliminated", funcName)
	}
}

func TestSliceDeleteStringEleByIndex(t *testing.T) {
	s := []string{"mark", "marry", "apple"}

}
