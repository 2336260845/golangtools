package datastructure

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMapReverseString2(t *testing.T) {
	funcName := "TestMapReverseString2"
	m := map[string]string{"first": "1", "second" : "2"}

	s := MapReverseString2(m)
	if len(s) != 2 {
		t.Errorf("Test %s error, error = %v", funcName, fmt.Sprintf("the len of revesemap is not right, len = %v", strconv.FormatInt(int64(len(s)), 10)))
		return
	}
	if s["1"] != "first" || s["2"] != "second" {
		t.Errorf("Test %s error, error = %v", funcName,  fmt.Sprintf("map is not right, map = %+v", s))
		return
	}

}

func TestMapReverseStringInt(t *testing.T) {
	funcName := "TestMapReverseStringInt"
	m := map[string]int{"first": 1, "second" : 2}

	s := MapReverseStringInt(m)
	if len(s) != 2 {
		t.Errorf("Test %s error, error = %v", funcName, fmt.Sprintf("the len of revesemap is not right, len = %v", strconv.FormatInt(int64(len(s)), 10)))
		return
	}
	if s[1] != "first" || s[2] != "second" {
		t.Errorf("Test %s error, error = %v", funcName,  fmt.Sprintf("map is not right, map = %+v", s))
		return
	}

}

func TestMapReverseIntString(t *testing.T) {
	funcName := "TestMapReverseIntString"
	m := map[int]string{1: "first", 2 : "second"}

	s := MapReverseIntString(m)
	if len(s) != 2 {
		t.Errorf("Test %s error, error = %v", funcName, fmt.Sprintf("the len of revesemap is not right, len = %v", strconv.FormatInt(int64(len(s)), 10)))
		return
	}
	if s["first"] != 1 || s["second"] != 2 {
		t.Errorf("Test %s error, error = %v", funcName,  fmt.Sprintf("map is not right, map = %+v", s))
		return
	}

}
