package greet

import "testing"

func TestGreetAcceptsInputFromUser(t *testing.T) {
	t.Parallel()
	want := "Terrence"
	got := greet()

	if want != got {
		t.Fatalf("wanted %q got %q", want ,got)
	}
}
