package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 2)
	if result != 3 {
		t.Error("unexpected result")
	}
}
