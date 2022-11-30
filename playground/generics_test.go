package playground

import (
	"testing"
)

func TestTryGenericsWithInt(t *testing.T) {
	got := TryGenerics(1)
	want := 1

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, 1)
	}
}

func TestTryGenericsWithString(t *testing.T) {
	got := TryGenerics("1")
	want := "1"

	if got != want {
		t.Errorf("got %s want %s given, %s", got, want, "1")
	}
}
