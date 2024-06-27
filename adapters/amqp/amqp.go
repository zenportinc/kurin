package amqp

import (
	"context"
	"os"

	"github.com/assembla/cony"
	"github.com/streadway/amqp"
	"github.com/zenportinc/kurin"
	"github.com/zenportinc/service-core/go/tracing"
)

var tracer = tracing.NewTracer()

type (
	Adapter struct {
		client   *cony.Client
		consumer *cony.Consumer
		handler  DeliveryHandler
		onStop   chan os.Signal
		logger   kurin.Logger
	}

	DeliveryHandler func(msg amqp.Delivery)
)

func NewAMQPAdapter(client *cony.Client, consumer *cony.Consumer, handler DeliveryHandler, logger kurin.Logger) kurin.Adapter {
	return &Adapter{
		client:   client,
		consumer: consumer,
		handler:  tracingMiddleware(handler),
		logger:   logger,
	}
}

func (adapter *Adapter) Open() {
	adapter.logger.Info("Consuming amqp...")
	for adapter.client.Loop() {
		select {
		case msg := <-adapter.consumer.Deliveries():
			adapter.handler(msg)
		case err := <-adapter.client.Errors():
			adapter.logger.Fatal(err)
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

func tracingMiddleware(next DeliveryHandler) DeliveryHandler {
	return func(msg amqp.Delivery) {
		// TODO: Add context to the DeliveryHandler
		_, span := tracer.Start(context.TODO(), tracing.SpanName())
		defer span.End()

		next(msg)
	}
}
