package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kafka-poc-consumer/internal/events"
	"log"
	"sync"
	"time"

	geutils "github.com/GenaroDaniel/geutils/pkg/events"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInterface interface {
	Produce(ctx context.Context, topic string, body any) error
	Consume(ctx context.Context)
}

type Kafka struct {
	Client     *kgo.Client
	Dispatcher geutils.EventDispatcher
}

type KafkaMessage struct {
	Key       string
	Topic     string
	Value     []byte
	Timestamp time.Time
}

func NewKafka(seeds []string, topics []string, dispatcher geutils.EventDispatcher) (KafkaInterface, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %s", err.Error())
	}

	return &Kafka{
		Client:     client,
		Dispatcher: dispatcher,
	}, nil
}

func (k *Kafka) Produce(ctx context.Context, topic string, body any) error {
	payload, err := k.SerializePayload(body)
	if err != nil {
		return err
	}

	record := &kgo.Record{Topic: topic, Value: payload}
	k.Client.ProduceSync(ctx, record)
	return nil
}

func (k *Kafka) SerializePayload(body any) ([]byte, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Error while serializing payload %s", err.Error())
	}
	return response, nil
}

func (k *Kafka) Consume(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down...")
			k.Client.Close()
			return
		default:
			fetches := k.Client.PollFetches(ctx)
			var wg sync.WaitGroup
			fetches.EachPartition(func(p kgo.FetchTopicPartition) {
				wg.Add(1)
				func(p kgo.FetchTopicPartition) {
					defer wg.Done()
					for _, record := range p.Records {

						key := string(record.Key)

						switch key {
						case string(events.CreateOrder):
							event := &events.CreateOrderEvent{
								DateTime: time.Now(),
								Payload:  record.Value,
							}
							fmt.Println("Received createOrder event:", string(record.Value))
							err := k.Dispatcher.Dispatch(ctx, event)
							if err != nil {
								fmt.Println("createOrder event error:", err)
							}

						default:
							fmt.Println("Unknown event key:", key)
						}

					}
				}(p)
			})
			wg.Wait()
		}
	}
}
