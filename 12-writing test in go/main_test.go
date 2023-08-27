package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 100.0, 20.0, 5.0, false},
}

// setiap test function harus diawali Test perhatikan kapital awalnya
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have ")
	}
}
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Did not give an error when we should have error cant divide with 0")
	}
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but we dont get error", "<<", tt.name)
			}
		} else {
			if err != nil {
				t.Error("did not expected an error but we got one", "<<", tt.name)
			}
		}

		if got != tt.expected {
			t.Error("expected error cause ", got, "is not equal as", tt.expected, "<<", tt.name)
		}

	}
}
