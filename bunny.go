package bunny

type Client struct {
	URI string
}

func New() (Client, error) {
	return Client{
			URI: "amqp://guest:guest@localhost/%2f",
		},
		nil
}
