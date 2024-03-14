package main

import "testing"

func TestAdd(t *testing.T) {
	// Test Driven Development (TDD)
	t.Run("add positive number", func(t *testing.T) {
		res := add(1, 2)
		if res != 3 {
			t.Errorf("failed to add 1 + 2 got %v expected 3", res)
		}
	})

	t.Run("add very big positive number positive number", func(t *testing.T) {
		res := add(1, 2)
		if res != 3 {
			t.Errorf("failed to add 1 + 2 got %v expected 3", res)
		}
		if res == 3 {
			t.Log("ok")
		}
	})

	t.Run("add positive and negative number", func(t *testing.T) {
		res := add(1, -2)
		if res != -1 {
			t.Errorf("failed to add 1 + (-2) got %v expected -1", res)
		}
	})

	t.Run("add negative negative number", func(t *testing.T) {
		res := add(-1, -2)
		if res != -3 {
			t.Errorf("failed to add 1 + (-2) got %v expected -3", res)
		}
	})
}
