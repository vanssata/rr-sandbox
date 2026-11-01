package greet

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("rr"); got != "Hello, rr" {
		t.Fatalf("Hello(rr) = %q", got)
	}
}
