package datastructure

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMapReverseString2(t *testing.T) {
	funcName := "TestMapReverse"
	m := map[string]string{"first": "1", "second" : "2"}

	s := MapReverseString2(m)
	if len(s) != 2 {
		t.Errorf("Test %s error, error = %v", funcName, "the len of revesemap is not right, len = " + strconv.FormatInt(int64(len(s)), 10))
		return
	}
	if s["1"] != "first" || s["2"] != "second" {
		t.Errorf("Test %s error, error = %v", funcName, "map is not right, map = " + fmt.Sprintf("%+v", s))
		return
	}

}
