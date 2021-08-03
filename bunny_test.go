package bunny_test

import (
	"bunny"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBunny(t *testing.T) {
	t.Parallel()
	var got bunny.Client
	got, err := bunny.New()
	if err != nil {
		t.Fatal(err)
	}

	want := bunny.Client{
		URI: "amqp://guest:guest@localhost/%2f",
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
