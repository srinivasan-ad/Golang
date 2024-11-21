package producer
import(
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"encoding/json"
	"log"
)
type ProducerConfig struct {
	Broker string
	Topic  string
}
func KafkaProducer(config ProducerConfig ) error {
	ctx := context.Background()
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.Broker},
		Topic: config.Topic,
	})
	defer kafkaWriter.Close()

	records, err := ReadCSV()
	if err != nil {
		log.Printf("Failed to read CSV: %v", err)
		return err
	}
	for _, record := range records {
		dataMap := make(map[string]string)
		for i, value := range record {
			dataMap[fmt.Sprintf("Column%d", i+1)] = value
		}

		message, err := json.Marshal(dataMap)
		if err != nil {
			log.Printf("Error marshalling data into JSON: %v", err)
			continue
		}

		err = kafkaWriter.WriteMessages(ctx, kafka.Message{
			Value: message,
		})
		if err != nil {
			log.Printf("Error sending message to Kafka: %v", err)
			continue
		}

		log.Printf("Sent message to Kafka: %s", message)
	}

	kafkaWriter.Close()
	return nil
}

