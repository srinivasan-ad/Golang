package consumer
import (
	"context"
	"fmt"
	"log"
    "github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
)
type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
func kafkaConsumer(){
	ctx := context.Background()
	const connectionString = "postgresql://postgres.ejspmxevlxzxowxttbmy:Vilasbhai94!@aws-0-ap-south-1.pooler.supabase.com:6543/postgres"
	pool,err := pgxpool.New(ctx,connectionString)
	if err != nil{
		log.Fatalf("Error while connecting to supabase :( %v",err)
	}
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "test-topic",
		GroupID: "consumer-group",
       //You can limit the bytes of data fetched using minByte and maxByte if needed
	})
	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Error reading message from Kafka: %v", err)
		}
		fmt.Printf("Received message from Kafka: %s\n", string(msg.Value))

		var firstName, lastName string
		var FirstName, LastName string
		var id int
		_, err = fmt.Sscanf(string(msg.Value), "First Name: %s, Last Name: %s", &firstName, &lastName)
		if err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}
		query := "INSERT INTO names (firstName, lastName) VALUES ($1, $2) RETURNUNG *"
		err = pool.QueryRow(ctx, query, firstName, lastName).Scan(&id, &FirstName, &LastName)
		if err != nil {
			log.Printf("Failed to insert data into DB: %v", err)
			} else 
			{
				fmt.Println("Inserted a new tuple successfully :)")
				fmt.Printf("ID: %d, FirstName: %s, LastName: %s\n", id, FirstName, LastName)
			}
			}
}

