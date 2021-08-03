package bunny_test

import (
	"bunny"
	"fmt"
	"testing"
)

func TestNewBunny(t *testing.T) {
	t.Parallel()
	var client bunny.Client
	client, err := bunny.New()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client)
}
