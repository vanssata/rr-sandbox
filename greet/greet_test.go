package greet

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("rr"); got != "Hello, rr" {
		t.Fatalf("Hello(rr) = %q", got)
	}
}

func TestShout(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"name", "rr", "HELLO, RR!"},
		{"empty", "", "HELLO, !"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Shout(tc.in); got != tc.want {
				t.Fatalf("Shout(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
