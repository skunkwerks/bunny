package bunny

import (
	"context"
	"fmt"
	"time"

	"github.com/makasim/amqpextra"
	"github.com/makasim/amqpextra/publisher"
	"github.com/streadway/amqp"
)

func (b Broker) Publish() error {
	// open connection
	conn, err := amqpextra.NewDialer(amqpextra.WithURL(b.URI))
	if err != nil {
		return fmt.Errorf("unable to establish connection to %v: %w", b.URI, err)
	}
	// close connection gracefully on termination
	defer conn.Close()

	// create publisher
	pub, err := conn.Publisher()
	if err != nil {
		return fmt.Errorf("unable to create publisher for %v: %w", b.URI, err)
	}
	// close publisher gracefully on termination
	defer pub.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	// publish a message
	for i := 0; i < 1_000_000; i++ {
		body := fmt.Sprintf("{\"index\": %v}", i)
		err = pub.Publish(publisher.Message{
			Key:     "test",
			Context: ctx,
			Publishing: amqp.Publishing{

				Body: []byte(body),
			},
		})
		if err != nil {
			return fmt.Errorf("unable to publish message %v to broker %v: %w", i, b.URI, err)
		}
	}
	return nil
}
