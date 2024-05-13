package notification

import (
	"fmt"
	"os"

	"whale-hotel/config"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type IService interface {
	NotificationRegister()
}

type Service struct {
	config config.IService
}

func NewService() *Service {
	return &Service{
		config: config.NewService(),
	}
}

func (s Service) NotificationRegister() {
	conf := s.config.ReadConfig()
	// sets the consumer group ID and offset
	conf["group.id"] = "go-group-1"
	conf["auto.offset.reset"] = "earliest"
	topic := "whale-hotel-registry"

	// creates a new consumer and subscribes to your topic
	consumer, _ := kafka.NewConsumer(&conf)
	consumer.SubscribeTopics([]string{topic}, nil)

	run := true
	for run {
		// consumes messages from the subscribed topic and prints them to the console
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			run = false
		}
	}

	// closes the consumer connection
	consumer.Close()
}
