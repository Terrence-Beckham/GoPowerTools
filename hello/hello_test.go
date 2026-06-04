package hello_test

import (
	"hello"
	"testing"
)

func TestPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	hello.Print()
}
