package github.com/copter01252/go-module_test

import (
	"demo"
	"testing"
)

func TestSayHiWihtKrit(t *testing.T) {
	r := demo.SayHi("Krit")
	if r != "Say Hi Krit" {
		t.Fatalf("Error %v", r)
	}
}

func TestSayHiWihtEmptyString(t *testing.T) {
	r := demo.SayHi("")
	if r != "Error" {
		t.Fatalf("Error %v", r)
	}
}
