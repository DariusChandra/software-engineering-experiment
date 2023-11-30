package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataImporter struct {
	csvReader       *csv.Reader
	mongoCollection *mongo.Collection
}

func main() {
	// Open the CSV file
	fileLocation := "golang/csv_to_mongo/test.csv"
	csvFile, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a CSV reader
	csvReader := csv.NewReader(csvFile)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	mongoCollection := client.Database("mydatabase").Collection("mycollection")
	dataImporter := NewDataImporter(csvReader, mongoCollection)

	// Import the data
	failedRecords, err := dataImporter.ImportData()
	if err != nil {
		log.Println(err)
	}

	// write the failed records
	file, err := os.Create("golang/csv_to_mongo/records.csv")
	defer file.Close()
	if err != nil {
		fmt.Println("failed to open file", err)
	}
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.WriteAll(failedRecords)
	if err != nil {
		fmt.Println("failed to write the error to file : ", err)
	}
}

func NewDataImporter(csvReader *csv.Reader, mongoCollection *mongo.Collection) *DataImporter {
	return &DataImporter{
		csvReader:       csvReader,
		mongoCollection: mongoCollection,
	}
}

func (di *DataImporter) ImportData() (failedRecords [][]string, err error) {
	columnHeaders, err := di.csvReader.Read()
	if err != nil {
		return nil, err
	}
	fmt.Println("columnHeaders", columnHeaders)
	failedRecords = append(failedRecords, columnHeaders)
	for {
		record, err := di.csvReader.Read()
		if err != nil {
			//if err == csv.ErrTooManyFields {
			//	log.Println("Too many fields in CSV record, skipping...")
			//	continue
			//}

			if err == io.EOF {
				break // Reached the end of the CSV file
			}

			continue
		}
		fmt.Println("record", record)

		filter := bson.D{}
		update := bson.M{}
		for i := 0; i < len(columnHeaders); i++ {
			//update = append(update, bson.E{
			//	Key:   columnHeaders[i],
			//	Value: record[i],
			//})
			update[columnHeaders[i]] = record[i]
			if columnHeaders[i] == "_id" {
				filter = append(filter, bson.E{
					Key:   columnHeaders[i],
					Value: record[i],
				})
			}
		}
		//updateMongo := bson.M{"$setOnInsert": update}
		updateMongo := bson.M{"$setOnInsert": update}
		updateOpts := options.Update().SetUpsert(true)
		result, err := di.mongoCollection.UpdateOne(context.Background(), filter, updateMongo, updateOpts)
		if err != nil {
			failedRecords = append(failedRecords, record)
			fmt.Println("failed to write to mongodb, the error", err)
			continue
		}
		if result.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document")
		}
		if result.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		}
	}

	fmt.Println("Finished imported data from CSV to MongoDB")
	return failedRecords, nil
}

func createDate(year, month, day int32) time.Time {
	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
}
