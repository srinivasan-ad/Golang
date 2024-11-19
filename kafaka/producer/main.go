package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main(){
	ctx := context.Background() 
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"} ,
		Topic : "test-topic",
	})
	file,err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("Could not open the CSV file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records,err := reader.ReadAll()
	if err != nil{
		log.Fatalf("Error while reading file %v",err)
	}
	for _,record := range records{
		message := fmt.Sprintf("First Name: %s, Last Name: %s", record[0], record[1]) 
		err := kafkaWriter.WriteMessages(ctx,kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			log.Fatalf("Error writing message to Kafka: %v", err)
		}
		fmt.Println("Sent to Kafka:", message)
	}
	defer kafkaWriter.Close()
}