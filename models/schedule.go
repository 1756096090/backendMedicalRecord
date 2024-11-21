package models


import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Date              time.Time          `bson:"date"`
	IDUser            string             `bson:"id_user"`
	IDPatient         string             `bson:"id_patient"`
	StartAppointment  time.Time          `bson:"start_appointment"`
	EndAppointment    time.Time          `bson:"end_appointment"`
	StartOriginalDate *time.Time         `bson:"start_original_date,omitempty"`
	Text              string             `bson:"text"`
}
