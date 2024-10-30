package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Specialist struct {
    ID      primitive.ObjectID `bson:"_id,omitempty"`
    Specialization     string `bson:"specialization"`
    Description        string `bson:"description"`
    YearsOfExperience  int    `bson:"yearsOfExperience"`
}
