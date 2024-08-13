package main

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-mongodb/v3/dplock"
	mongoDriver "github.com/ONSdigital/dp-mongodb/v3/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"sync"
)

func main() {
	// Create a Mongo session and set the write mode to "majority".
	mongoUrl := "mongodb://localhost:27017"
	database := "adtech-adcredits"
	resource := "adcredits"
	ctx := context.Background()
	m, err := mongo.Connect(ctx, options.Client().
		ApplyURI(mongoUrl).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority())))

	if err != nil {
		log.Fatal(err)
	}
	mongoConnection := mongoDriver.NewMongoConnection(m, database)
	lock := dplock.New(ctx, mongoConnection, resource)

	balance := 100
	var wg sync.WaitGroup
	for i := 0; i < balance; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := "abcd1234"
			lockID, err := lock.Acquire(ctx, id)
			defer lock.Unlock(ctx, lockID)
			if err != nil {
				fmt.Println("got error lock", err)
				return
			}
			balance--
			fmt.Println(balance)
		}()
	}
	wg.Wait()
	fmt.Println("balance", balance)
}
