package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Patient struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Name                     string             `bson:"name"`
	BirthDate                time.Time          `bson:"birthDate"`
	Gender                   bool               `bson:"gender"`
	Concerning               string             `bson:"concerning"`
	Mail                     string             `bson:"mail"`
	DNI                      string             `bson:"dni"`
	Phone                    string             `bson:"phone"`
	Occupation               string             `bson:"occupation"`
	Responsible              string             `bson:"responsible"`
	HasInsurance             bool               `bson:"hasInsurance"`
	HasHeartDisease          bool               `bson:"hasHeartDisease"`
	HasBloodPressure         bool               `bson:"hasBloodPressure"`
	HasBloodGlucose          bool               `bson:"hasBloodGlucose"`
	HasDiabetes              bool               `bson:"hasDiabetes"`
	HasAllergies             bool               `bson:"hasAllergies"`
	HasEndocrineDisorders    bool               `bson:"hasEndocrineDisorders"`
	HasNeurologicalDisorders bool               `bson:"hasNeurologicalDisorders"`
	Others                   string             `bson:"others"`
}
