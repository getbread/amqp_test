package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getbread/amqp_test_app/shared"
	"github.com/getbread/desmond"
	"github.com/getbread/desmond/request"
)

func createDesmondConnection(serviceName string) *desmond.Queue {
	config := desmond.AmqpConfig{
		User:     os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		AppId:    serviceName,
	}
	config.DefaultTimeout = 15 * time.Second
	return desmond.NewQueue(desmond.NewAmqpConnection(config))
}

func main() {
	shared.LoadEnvironment()

	requester := request.NewRequester(createDesmondConnection("METRICS_SUBSCRIBER"))

	requester.Handle(request.CreateMetricA.Request(), func(c *request.RequestContext) {
		var req shared.Payload
		err := c.Bind(&req)
		if err != nil {
			panic(err)
		}

		c.Respond(req)
	})

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	println("Shutting down...")
	return

}
