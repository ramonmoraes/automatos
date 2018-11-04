package models

import "testing"

func TestGetSimbol(t *testing.T) {
	s1, err := NewSimbol("r")
	if err != nil || s1.IsVariable == true {
		t.Error("Expected 'r' to not be a variable and not have errors")
	}

	s2, err := NewSimbol("R")
	if err != nil || s2.IsVariable != true {
		t.Error("Expected 'R' to be variable and not have errors")
	}

	s3, err := NewSimbol("Rar")
	if err == nil || s3.Value != "" {
		t.Error("Expected 'Rar' to have error")
	}
}
