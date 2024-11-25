package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Procedure struct {
    ID      primitive.ObjectID `bson:"_id,omitempty"`
    Description        string `bson:"description"`
    IsTimeType  bool    `bson:"is_time_type"`
	
}
