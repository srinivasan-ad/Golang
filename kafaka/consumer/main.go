package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	connectionString := "postgresql://postgres.ejspmxevlxzxowxttbmy:Vilasbhai94!@aws-0-ap-south-1.pooler.supabase.com:6543/postgres"

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}
	defer pool.Close()

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "consumer-group",
	})
	defer kafkaReader.Close()

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Error reading message from Kafka: %v", err)
		}

		fmt.Printf("Received message from Kafka: %s\n", string(msg.Value))

		var firstName, lastName string
		msgString := string(msg.Value)

		parts := strings.Split(msgString, ", ")
		if len(parts) == 2 {
			firstName = strings.TrimPrefix(parts[0], "First Name: ")
			lastName = strings.TrimPrefix(parts[1], "Last Name: ")
		} else {
			log.Printf("Unexpected message format: %s", msgString)
			continue
		}

		var id int
		query := "INSERT INTO names (firstName, lastName) VALUES ($1, $2) RETURNING id"
		err = pool.QueryRow(ctx, query, firstName, lastName).Scan(&id)
		if err != nil {
			log.Printf("Failed to insert data into DB: %v", err)
		} else {
			fmt.Println("Inserted a new tuple successfully :)")
			fmt.Printf("ID: %d, FirstName: %s, LastName: %s\n", id, firstName, lastName)
		}
	}
}
