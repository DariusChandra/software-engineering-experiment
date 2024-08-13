package main

import (
	"context"
	lock "github.com/square/mongo-lock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"time"
)

func main() {
	mongoUrl := "mongodb://localhost:27017"
	database := "adtech-adcredits"
	collection := "distributed-locks"

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

	lockId := "abcd1234"

	// Create an exclusive lock on resource1.
	err = c.XLock(ctx, "resource1", lockId, lock.LockDetails{
		TTL: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a shared lock on resource2.
	err = c.SLock(ctx, "resource2", lockId, lock.LockDetails{
		Owner:   "",
		Host:    "",
		Comment: "",
		TTL:     5,
	}, -1)
	if err != nil {
		log.Fatal(err)
	}

	//// Unlock all locks that have our lockId.
	//_, err = c.Unlock(ctx, lockId)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
