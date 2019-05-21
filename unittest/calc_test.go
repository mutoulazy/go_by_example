package main 

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add r error %d", r)
	}
	t.Logf("add success")
}

func TestSub(t *testing.T) {
	r := sub(4, 2)
	if r != 2 {
		t.Fatalf("sub r error %d", r)
	}
	t.Logf("sub success")
}