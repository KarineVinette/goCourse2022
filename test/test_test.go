package main

import "testing"

func TestIntMin(t *testing.T) {
	ans := IntMin(12, 15)
	if ans != 12 {

		t.Errorf("IntMin(12, 15) = %d; want 12", ans)
	}
}
