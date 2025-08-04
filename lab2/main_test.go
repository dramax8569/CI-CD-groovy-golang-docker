package main

import "testing"

func TestPing(t *testing.T) {
	expected := "pong"
	result := Ping()
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
