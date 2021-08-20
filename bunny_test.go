package bunny_test

import (
	"bunny"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBunnyWithDefaultURI(t *testing.T) {
	t.Parallel()
	var got bunny.Broker
	got, err := bunny.NewBroker()
	if err != nil {
		t.Fatal(err)
	}

	want := bunny.Broker{
		URI: "amqp://guest:guest@localhost/",
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestNewBunnyWithCustomURI(t *testing.T) {
	t.Parallel()
	custom_uri := "amqps://wascally:wabbit@localhost:port/vhost"

	var got bunny.Broker
	got, err := bunny.NewBrokerFromArgs([]string{"-u", custom_uri})
	if err != nil {
		t.Fatal(err)
	}

	want := bunny.Broker{
		URI: custom_uri,
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
