package main

import "testing"

func TestShake(t *testing.T) {
	// test to make sure strings yielded
	result := Shake()
	if len(result) < 1 {
		t.Error("Test yielded empty string")
	}
}
