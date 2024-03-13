package amqp

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zenportinc/carrotmq"
	"github.com/zenportinc/kurin"
)

type (
	Adapter struct {
		client   *carrotmq.Client
		consumer *carrotmq.Consumer
		handler  DeliveryHandler
		onStop   chan os.Signal
		logger   kurin.Logger
	}

	DeliveryHandler func(msg amqp.Delivery)
)

func NewAMQPAdapter(client *carrotmq.Client, consumer *carrotmq.Consumer, handler DeliveryHandler, logger kurin.Logger) kurin.Adapter {
	return &Adapter{
		client:   client,
		consumer: consumer,
		handler:  handler,
		logger:   logger,
	}
}

func (adapter *Adapter) Open() {
	adapter.logger.Info("Consuming amqp... test")
	adapter.client.StartClient()
	adapter.client.Consume()
	for {
		select {
		case msg := <-adapter.consumer.Deliveries():
			adapter.handler(msg)
		case err := <-adapter.client.Errors():
			adapter.logger.Error("Client error: %v\n", err)
		case err := <-adapter.consumer.Errors():
			adapter.logger.Error("Consumer error: %v\n", err)
		}

	}
}

func (adapter *Adapter) Close() {
	adapter.client.Close()
}

func (adapter *Adapter) OnFailure(err error) {
	if err != nil {
		adapter.logger.Fatal(err)
	}
}
