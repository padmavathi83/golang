package main

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(12, 3)
	if ans != 10 {
		t.Error("add was incorrect, got: %d, want:%d", ans, 10)
	}
}
