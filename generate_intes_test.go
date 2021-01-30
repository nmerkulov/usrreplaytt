package main

import (
	"testing"
)

func Test_generateInts(t *testing.T) {
	exp := 42
	sampleFunc := func() int {
		return exp
	}
	ch := generateInts(1, sampleFunc)

	if got := <-ch; got != exp {
		t.Errorf("expected :%d; got: %d", exp, got)
	}

	if _, ok := <-ch; ok {
		t.Errorf("expected channel to be closed but it is open")
	}
}
