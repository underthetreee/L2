package main

import "testing"

func TestOr(t *testing.T) {
	done1 := make(chan interface{})
	done2 := make(chan interface{})
	done3 := make(chan interface{})

	single := or(done1, done2, done3)

	close(done2)

	_, ok := <-single
	if !ok {
		t.Log("single channel closed as expected")
	} else {
		t.Error("single channel still open")
	}
}
