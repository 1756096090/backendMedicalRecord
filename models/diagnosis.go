package models

type Diagnosis struct {
	ID             string           `bson:"_id,omitempty"`
	Code           string           `bson:"code"`
	Description    string           `bson:"name"`
	UnderDiagnosis []UnderDiagnosis `bson:"under_diagnosis"`
}

type UnderDiagnosis struct {
	Code        string `bson:"code"`
	Description string `bson:"name"`
}
