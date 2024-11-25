package models

import (
	"time"
)

type Schedule struct {
	ID               string     `bson:"_id"`
	Date             time.Time  `bson:"date"`
	IDUser           string     `bson:"id_user"`
	IDPatient        string     `bson:"id_patient"`
	StartAppointment time.Time  `bson:"start_appointment"`
	EndAppointment   time.Time  `bson:"end_appointment"`
	StartOriginal    *time.Time `bson:"start_original_date,omitempty"`
	Text             string     `bson:"text"`
}

type ScheduleDetail struct {
	Schedule `bson:",inline"` 
	User     User     `bson:"user"`
	Patient  Patient  `bson:"patient"`
}