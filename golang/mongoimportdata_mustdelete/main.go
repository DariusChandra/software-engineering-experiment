package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	// MongoDB connection URI
	uri := "mongodb://localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Define the collection
	collection := client.Database("adtech-advertisers").Collection("advertisers")

	isActive := true
	fileName := "vendor_ids_active.txt"
	if !isActive {
		fileName = "vendor_ids_inactive.txt"
	}

	// Define the filter
	filter := bson.M{"global_entity_id": "FP_SG", "_object_type": "vendor", "active": isActive}

	// Find documents
	// Set the find options
	findOptions := options.Find()
	findOptions.SetBatchSize(100)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Open the file for appending
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Track if this is the first write for comma separation
	isFirst := true

	// Iterate through the cursor
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// Assuming 'vendorID' is the field we want to extract
		if vendorID, ok := result["vendor_id"].(string); ok {
			fmt.Println(vendorID)
			// Write the vendorID to the file
			if isFirst {
				_, err = file.WriteString(vendorID)
				isFirst = false
			} else {
				_, err = file.WriteString("," + vendorID)
			}

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor
	cursor.Close(context.TODO())

	fmt.Printf("Vendor IDs written to %s \n", fileName)

	// Disconnect from MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
