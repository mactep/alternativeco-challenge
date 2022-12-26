package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mactep/alternativeco-challenge/email/ent"

	_ "github.com/lib/pq"
)

func main() {
	client := Connect()
	defer client.Close()
	ctx := context.Background()

	// Create email
	email, err := CreateEmail(ctx, client, "narutinho@vrau.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)

	// Delete email
	err = DeleteEmail(ctx, client, email.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func Connect() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=email sslmode=disable password=password")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("Connected to database")
	return client
}

func CreateEmail(ctx context.Context, client *ent.Client, email string) (*ent.Email, error) {
	e, err := client.Email.Create().SetEmail(email).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating email: %w", err)
	}
	log.Println("email created: ", e)
	return e, nil
}

func DeleteEmail(ctx context.Context, client *ent.Client, id int) error {
	err := client.Email.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting email: %w", err)
	}
	log.Println("email deleted: ", id)
	return nil
}
