package greet

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "rr", "Hello, rr"},
		{"empty", "", "Hello, "},
		{"spaces", "John Smith", "Hello, John Smith"},
		{"non-ascii", "Иван", "Hello, Иван"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.in); got != tt.want {
				t.Fatalf("Hello(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
