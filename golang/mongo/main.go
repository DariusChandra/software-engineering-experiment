package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"regexp"
	"time"
)

func main() {
	pipelineUpdate()

	//filter := createSearchFilter("b066", "Expired", "FP_SG")
	//filterByte, err := bson.MarshalExtJSON(filter, true, false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(filterByte))
}

func pipelineUpdate() {
	// Set up a MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Select your database and collection
	collection := client.Database("adtech-adcredits").Collection("adcredits")

	objID, err := primitive.ObjectIDFromHex("656ec3e1b8cef77f8b2545a1")
	filter := bson.M{
		"_id": objID,
	}
	utilisationFloat := 90.00
	setStage := bson.M{
		"utilization": bson.M{"$max": bson.A{"$utilization", utilisationFloat}},
		// Set 'is_consumed' based on a condition
		"is_consumed": bson.M{
			"$cond": bson.A{
				bson.M{"$gt": bson.A{utilisationFloat, "$utilization"}}, // Condition: if incomingUtilization is greater
				bson.M{"$eq": bson.A{"$credit", utilisationFloat}},      // Then: check if incomingUtilization equals credit
				"$is_consumed", // Else: keep the current value of is_consumed
			},
		},
		"balance":     bson.M{"$min": bson.A{"$balance",bson.M{"$subtract": bson.A{"$credit", utilisationFloat}}}},
	}
	setStage["last_utilization"] = bson.M{"$max": bson.A{"$last_utilization", time.Now().Add(24 * time.Hour)}}
	update := bson.A{
		bson.M{
			"$set": setStage,
		},
	}

	collection.UpdateOne(ctx, filter, update)
}

func createSearchFilter(searchQuery string, status string, globalEntityId string) bson.M {
	filters := bson.M{"global_entity_id": globalEntityId}

	if searchQuery != "" {
		containsQuery := ".*" + regexp.QuoteMeta(searchQuery) + ".*"
		pattern := primitive.Regex{Pattern: containsQuery, Options: "i"}

		filters["$or"] = []bson.M{
			{"advertiser_id": pattern},
			{"advertiser_name": pattern},
			{"program_name": pattern},
		}
	}

	if status != "" {
		today := time.Now().UTC().Truncate(24 * time.Hour)
		switch status {
		case "Expired":
			filters["$and"] = []bson.M{
				{"end_date": bson.M{"$lt": today}},
			}
		case "New":
			filters["$and"] = []bson.M{
				{"end_date": bson.M{"$gte": today}},
				{"utilization": 0},
			}
		case "Active":
			filters["$and"] = []bson.M{
				{"end_date": bson.M{"$gte": today}},
				{"utilization": bson.M{"$gt": 0}},
				{"$expr": bson.M{"$lt": []interface{}{"$utilization", "$credit"}}},
			}
		case "Consumed":
			filters["$and"] = []bson.M{
				{"end_date": bson.M{"$gte": today}},
				{"$expr": bson.M{"$eq": []interface{}{"$utilization", "$credit"}}},
			}
		}
	}

	return filters
}
