package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AdCredits struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AdvertiserID   string             `json:"advertiser_id" bson:"advertiser_id"`
	AdvertiserName string             `json:"advertiser_name" bson:"advertiser_name"`
	ProgramName    string             `json:"program_name" bson:"program_name"`
	CreatedBy      string             `json:"created_by" bson:"created_by"`
	UpdatedBy      string             `json:"updated_by" bson:"updated_by"`
	StartDate      time.Time          `json:"start_date" bson:"start_date"`
	EndDate        time.Time          `json:"end_date" bson:"end_date"`
	Credit         float64            `json:"credit" bson:"credit"`
	Utilization    float64            `json:"utilization" bson:"utilization"`
	GlobalEntityId string             `json:"global_entity_id" bson:"global_entity_id"`
	Timezone       string             `json:"timezone" bson:"timezone"`
}

func main() {
	adcreditUpsert := AdCredits{}
	b, _ := json.Marshal(&adcreditUpsert)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}
