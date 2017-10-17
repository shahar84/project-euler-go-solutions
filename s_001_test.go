package main

import "testing"

func TestToAdd(t *testing.T) {
	result := toAdd(0)
	if result != 0 {
		t.Error("Test fail for number ", result)
	}
	result = toAdd(17)
	if result != 0 {
		t.Error("Test fail for number ", result)
	}
	result = toAdd(3)
	if result != 3 {
		t.Error("Test fail for number ", result)
	}
	result = toAdd(5)
	if result != 5 {
		t.Error("Test fail for number ", result)
	}
}
