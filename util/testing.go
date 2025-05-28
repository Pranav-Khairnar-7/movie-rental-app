package util

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	} else {
		fmt.Println("Assertion passed")
	}
}
