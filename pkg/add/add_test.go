package add

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Fatalf("1+1 should be 2")
	}
}
