package main

import (
	"context"
	"fmt"
	lock "github.com/square/mongo-lock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"time"
)

func main() {
	// Create a Mongo session and set the write mode to "majority".
	mongoUrl := "mongodb://localhost:27017"
	database := "adtech-adcredits"
	collection := "adcredits_locks"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	m, err := mongo.Connect(ctx, options.Client().
		ApplyURI(mongoUrl).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority())))

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = m.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Configure the client for the database and collection the lock will go into.
	col := m.Database(database).Collection(collection)

	// Create a MongoDB lock client.
	c := lock.NewClient(col)

	// Create the required and recommended indexes.
	c.CreateIndexes(ctx)

	p := lock.NewPurger(c)
	lockStatus, err := p.Purge(ctx)
	fmt.Println(lockStatus, err)
}
