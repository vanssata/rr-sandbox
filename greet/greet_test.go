package greet

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("rr"); got != "Hello, rr" {
		t.Fatalf("Hello(rr) = %q", got)
	}
}

func TestFarewell(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"ascii", "rr", "Goodbye, rr"},
		{"empty", "", "Goodbye, "},
		{"non-ascii", "Zoë 世界", "Goodbye, Zoë 世界"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Farewell(tc.in); got != tc.want {
				t.Fatalf("Farewell(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
