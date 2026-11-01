package greet

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	if got := Hello("rr"); got != "Hello, rr" {
		t.Fatalf("Hello(rr) = %q", got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("rr")
	}
}

func ExampleHello() {
	fmt.Println(Hello("rr"))
	// Output: Hello, rr
}
