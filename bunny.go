package bunny

import (
	"flag"
)

var default_amqp_uri = "amqp://guest:guest@localhost/"

type Broker struct {
	URI string
}

func NewBroker() (Broker, error) {
	return Broker{
		URI: default_amqp_uri,
	}, nil
}

func NewBrokerFromArgs(args []string) (Broker, error) {
	fset := flag.NewFlagSet("broker", flag.ContinueOnError)
	uri := fset.String("u", default_amqp_uri, "Set AMQP Broker URI")
	err := fset.Parse(args)
	if err != nil {
		return Broker{}, err
	}
	b := Broker{
		URI: *uri,
	}
	return b, nil
}
