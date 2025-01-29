package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var client *db.Client

func InitFirebase() {
	 // Load environment variables from .env file
	 err := godotenv.Load()
	 if err != nil {
		 log.Fatalf("Error loading .env file: %v", err)
	 }
	 
	ctx := context.Background()
    conf := &firebase.Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
    }
	opt := option.WithCredentialsFile("./serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
        log.Fatalln("Error initializing app:", err)
    }

	client, err = app.Database(ctx)
    if err != nil {
        log.Fatalln("Error initializing database client:", err)
    }

	// Example usage of the database client
    ref := client.NewRef("restricted_access/secret_document")
    var data map[string]interface{}
    if err := ref.Get(ctx, &data); err != nil {
        log.Fatalln("Error reading from database:", err)
    }
    fmt.Println(data)
}