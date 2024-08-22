package hello

import (
	"context"
	"testing"
)

func TestHello(t *testing.T) {
	resp, err := World(context.Background(), "Suraj")
	if err != nil {
		t.Fatal(err)
	}

	want, got := "Hello Suraj!", resp.Message
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
