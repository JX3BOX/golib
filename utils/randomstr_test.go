package utils

import "testing"

func TestRandString(t *testing.T) {

	for i := 0; i < 200000; i++ {
		got := RandString(4)
		t.Logf("RandString() = %v", got)
	}
}
