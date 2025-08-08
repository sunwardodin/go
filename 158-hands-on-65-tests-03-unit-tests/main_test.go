package main

import "testing"

func TestDoSquare(t *testing.T) {
	total := doSquare(2, 3, square)
	if total != 13 {
		t.Errorf("Square was incorrect, got: %v, wanted: %d", total, 13)
	}
}
