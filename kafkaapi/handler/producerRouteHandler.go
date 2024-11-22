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

func ProducerHandler(ctx *gofr.Context) (interface{}, error)  {
		var req ProduceRequest
		if err := ctx.Bind(&req); err != nil {
		
		
		}
		config := producer.ProducerConfig{
			Broker: req.Broker,
			Topic:  req.Topic,
		}
		go func() {
			if err := producer.KafkaProducer(ctx,config, req.CSVFile); err != nil {
				log.Printf("Failed to produce messages: %v", err)
			}
		}()
		return "Producer process started", nil

	}
