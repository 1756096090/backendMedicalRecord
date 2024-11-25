package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiagnosisProcedure struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IDPatient   string             `bson:"id_patient"`
	CodeDiagnosis string             `bson:"code_diagnosis"`
	CodeUnderDiagnosis string             `bson:"code_under_diagnosis"`
	Procedures  []ProcedureDetails `bson:"procedures"`
}

type ProcedureDetails struct {
	IDProcedure string     `bson:"id_procedure"`
	IDCreator   string     `bson:"id_creator"`
	IDUpdater   string     `bson:"id_updater"`
	Description string     `bson:"description"`
	StartAt     *time.Time `bson:"start_at,omitempty"`
	EndAt       *time.Time `bson:"end_at,omitempty"`
	IsCompleted bool       `bson:"is_completed"`
}
