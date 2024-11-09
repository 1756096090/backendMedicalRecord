// repository/patient_repository.go
package repository

import (
	"backendMedicalRecord/config"
	"backendMedicalRecord/models"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreatePatient(patient models.Patient) (*mongo.InsertOneResult, error) {
	patientCollection := config.DB.Collection("patient")
	log.Printf("Using collection: %s", patientCollection.Name())
	result, err := patientCollection.InsertOne(context.Background(), patient)
	if err != nil {
		log.Println("Error creating patient:", err)
		return nil, err
	}
	return result, nil
}

func GetPatientByID(patientID string) (*models.Patient, error) {
    patientCollection := config.DB.Collection("patient") 

    objID, err := primitive.ObjectIDFromHex(patientID)
    if err != nil {
        return nil, err 
    }

    var patient models.Patient
    err = patientCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&patient)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving patient:", err)
        return nil, err
    }
    
    return &patient, nil
}
func UpdatePatient(patientID string, updatedPatient models.Patient) (*mongo.UpdateResult, error) {
	patientCollection := config.DB.Collection("patient")
	log.Printf("Using collection: %s", patientCollection.Name())

	objID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedPatient}

	result, err := patientCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating patient:", err)
		return nil, err
	}
	return result, nil
}

func DeletePatient(patientID string) (*mongo.DeleteResult, error) {
	patientCollection := config.DB.Collection("patient")
	objID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := patientCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting patient:", err)
		return nil, err
	}
	return result, nil
}

func GetAllPatients() ([]models.Patient, error) {
	patientCollection := config.DB.Collection("patient")

	cursor, err := patientCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting patients:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var patients []models.Patient
	for cursor.Next(context.Background()) {
		var patient models.Patient
		if err := cursor.Decode(&patient); err != nil {
			log.Println("Error decoding patient:", err)
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return patients, nil
}


func GetPatientByDNIOrEmail(dni string, email string) (*models.Patient, error) {
    patientCollection := config.DB.Collection("patient")

    filter := bson.M{"$or": []bson.M{
        {"dni": dni},
        {"email": email},
    }}

    var patient models.Patient
    err := patientCollection.FindOne(context.Background(), filter).Decode(&patient)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }

    return &patient, nil
}




