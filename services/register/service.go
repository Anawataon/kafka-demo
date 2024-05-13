package register

import (
	"fmt"

	"whale-hotel/config"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type IService interface {
	ProduceRegister()
}

type Service struct {
	config config.IService
}

func NewService() *Service {
	return &Service{
		config: config.NewService(),
	}
}

func (s Service) ProduceRegister() {
	// creates a new producer instance
	conf := s.config.ReadConfig()
	p, _ := kafka.NewProducer(&conf)
	topic := "whale-hotel-registry"

	// go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// produces a sample message to the user-created topic
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("key"),
		Value: []byte(`{
			"firstName": "John",
			"lastName": "Doe",
			"roomNo2": 102,
		}`),
	}, nil)

	// send any outstanding or buffered messages to the Kafka broker and close the connection
	p.Flush(15 * 1000)
	p.Close()
}
