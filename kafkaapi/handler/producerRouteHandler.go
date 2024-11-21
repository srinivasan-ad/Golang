package handler
import(
	"gofr-server/kafkaapi/producer"
	"log"
	"gofr.dev/pkg/gofr"
)

type ProduceRequest struct {
	Broker  string `json:"broker"`
	Topic   string `json:"topic"`
	CSVFile string `json:"csvFile"`
}

func ProducerHandler(c *gofr.Context)  {
		var req ProduceRequest
		if err := c.Bind(&req); err != nil {
		
		
		}
		config := producer.ProducerConfig{
			Broker: req.Broker,
			Topic:  req.Topic,
		}
		go func() {
			if err := producer.KafkaProducer(config, req.CSVFile); err != nil {
				log.Printf("Failed to produce messages: %v", err)
			}
		}()

	}
