package hello_test

import (
	"hello"
	"testing"
)

func TestPrintPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	hello.Print()

}
