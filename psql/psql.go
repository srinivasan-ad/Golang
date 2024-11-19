package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)	

func main(){
	firstName := "Rama"
	lastName := "Krishna"
	connectionString := "postgresql://postgres.ejspmxevlxzxowxttbmy:Vilasbhai94!@aws-0-ap-south-1.pooler.supabase.com:6543/postgres"
	 ctx := context.Background()

	pool,err := pgxpool.New(ctx,connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Connected to database :)")
	var id int
	var insertedFirstName, insertedLastName string
	query := "INSERT INTO names (firstName, lastName) VALUES ($1, $2) RETURNING id, firstName, lastName"
	err = pool.QueryRow(ctx, query, firstName, lastName).Scan(&id, &insertedFirstName, &insertedLastName)

	if err != nil {
		log.Fatal("Error while executing query :(")
	}
	defer pool.Close()
	fmt.Println("Inserted a new tuple successfully :)")
	fmt.Printf("ID: %d, FirstName: %s, LastName: %s\n", id, insertedFirstName, insertedLastName)
}


