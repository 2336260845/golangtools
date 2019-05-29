package datastructure

import (
	"testing"
)
//-------------------------slice去重单测------------------------------------//
func TestSliceRemoveDupString(t *testing.T) {
	funcName := "TestSliceRemoveDupString"
	s := []string{"mark", "marry", "jack", "jack", "andy", "mark"}
	s, err := SliceRemoveDupString(s)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
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
		t.Errorf("Test %s error, error = %v", funcName, err)
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
		t.Errorf("Test %s error, error = %v", funcName, err)
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
		t.Errorf("Test %s error, error = %v", funcName, err)
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
		t.Errorf("Test %s error, error = %v", funcName, err)
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
//-------------------------slice删除元素单测------------------------------------//

func TestSliceDeleteStringEleByIndex(t *testing.T) {
	funcName := "TestSliceDeleteStringEleByIndex"
	s := []string{"mark", "marry", "apple"}
	s, err := SliceDeleteStringEleByIndex(s, 0)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
		return
	}
	if len(s) != 2 {
		t.Errorf("slice len is not right, len = %v, context = %+v", len(s), s)
		return
	}
	if s[0] != "marry" || s[1] != "apple" {
		t.Errorf("Test %s error, slice context is not right, context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteStringEleByIndex(s, len(s) - 1)
	if len(s) != 1 || s[0] != "marry" {
		t.Errorf("Test %v error, slice context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteStringEleByIndex(s, len(s))
	if err == nil {
		t.Errorf("Test %s error, out of index but not error", funcName)
		return
	}
}

func TestSliceDeleteIntEleByIndex(t *testing.T) {
	funcName := "TestSliceDeleteIntEleByIndex"
	s := []int{1, 2, 3}
	s, err := SliceDeleteIntEleByIndex(s, 0)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
		return
	}
	if len(s) != 2 {
		t.Errorf("slice len is not right, len = %v, context = %+v", len(s), s)
		return
	}
	if s[0] != 2 || s[1] != 3 {
		t.Errorf("Test %s error, slice context is not right, context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteIntEleByIndex(s, len(s) - 1)
	if len(s) != 1 || s[0] != 2 {
		t.Errorf("Test %v error, slice context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteIntEleByIndex(s, len(s))
	if err == nil {
		t.Errorf("Test %s error, out of index but not error", funcName)
		return
	}
}

func TestSliceDeleteInt64EleByIndex(t *testing.T) {
	funcName := "TestSliceDeleteInt64EleByIndex"
	s := []int64{1, 2, 3}
	s, err := SliceDeleteInt64EleByIndex(s, 0)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
		return
	}
	if len(s) != 2 {
		t.Errorf("slice len is not right, len = %v, context = %+v", len(s), s)
		return
	}
	if s[0] != 2 || s[1] != 3 {
		t.Errorf("Test %s error, slice context is not right, context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteInt64EleByIndex(s, len(s) - 1)
	if len(s) != 1 || s[0] != 2 {
		t.Errorf("Test %v error, slice context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteInt64EleByIndex(s, len(s))
	if err == nil {
		t.Errorf("Test %s error, out of index but not error", funcName)
		return
	}
}

func TestSliceDeleteInt32EleByIndex(t *testing.T) {
	funcName := "TestSliceDeleteInt32EleByIndex"
	s := []int32{1, 2, 3}
	s, err := SliceDeleteInt32EleByIndex(s, 0)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
		return
	}
	if len(s) != 2 {
		t.Errorf("slice len is not right, len = %v, context = %+v", len(s), s)
		return
	}
	if s[0] != 2 || s[1] != 3 {
		t.Errorf("Test %s error, slice context is not right, context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteInt32EleByIndex(s, len(s) - 1)
	if len(s) != 1 || s[0] != 2 {
		t.Errorf("Test %v error, slice context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteInt32EleByIndex(s, len(s))
	if err == nil {
		t.Errorf("Test %s error, out of index but not error", funcName)
		return
	}
}

func TestSliceDeleteByteEleByIndex(t *testing.T) {
	funcName := "TestSliceDeleteByteEleByIndex"
	s := []byte{'a', 'b', 'c'}
	s, err := SliceDeleteByteEleByIndex(s, 0)
	if err != nil {
		t.Errorf("Test %s error, error = %v", funcName, err)
		return
	}
	if len(s) != 2 {
		t.Errorf("slice len is not right, len = %v, context = %+v", len(s), s)
		return
	}
	if s[0] != 'b' || s[1] != 'c' {
		t.Errorf("Test %s error, slice context is not right, context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteByteEleByIndex(s, len(s) - 1)
	if len(s) != 1 || s[0] != 'b' {
		t.Errorf("Test %v error, slice context = %+v", funcName, s)
		return
	}
	s, err = SliceDeleteByteEleByIndex(s, len(s))
	if err == nil {
		t.Errorf("Test %s error, out of index but not error", funcName)
		return
	}
}

