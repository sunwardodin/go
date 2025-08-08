package main

import "testing"

func TestAdd(t *testing.T) {
	total := doMath(5, 5, add)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
