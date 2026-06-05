package hello_test

import (
	"bytes"
	"testing"
	"github.com/terrence-beckham/hello"
)

func TestPrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Fatalf("want %q, got %q",want, got)
	}
}
