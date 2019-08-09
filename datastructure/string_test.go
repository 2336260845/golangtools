package datastructure

import (
	"fmt"
	"testing"
)

func TestStringContainRuneOne(t *testing.T) {
	funcName := "TestStringContainRuneOne"
	s := "123456789"
	chyes := '2'
	chno := 'q'
	if y, i := StringContainRuneOne(s, chyes); !y || i != 1 {
		t.Errorf("Test %s error, error = %v", funcName, fmt.Sprintf("index or bool error: index = %d, bool = %v", i, y))
		return
	}

	if n, i := StringContainRuneOne(s, chno); n || i != -1 {
		t.Errorf("Test %s error, error = %v", funcName, fmt.Sprintf("index or bool error: index = %d, bool = %v", i, n))
		return
	}
}
