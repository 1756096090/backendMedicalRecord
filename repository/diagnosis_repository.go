// repository/diagnosis_repository.go
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

func CreateDiagnosis(diagnosis models.Diagnosis) (*mongo.InsertOneResult, error) {
	diagnosisCollection := config.DB.Collection("diagnosis")
	result, err := diagnosisCollection.InsertOne(context.Background(), diagnosis)
	if err != nil {
		log.Println("Error creating diagnosis:", err)
		return nil, err
	}
	return result, nil
}

func GetDiagnosisByID(diagnosisID string) (*models.Diagnosis, error) {
    diagnosisCollection := config.DB.Collection("diagnosis") 

    objID, err := primitive.ObjectIDFromHex(diagnosisID)
    if err != nil {
        return nil, err
    }

    var diagnosis models.Diagnosis
    err = diagnosisCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&diagnosis)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving diagnosis:", err)
        return nil, err
    }
    
    return &diagnosis, nil
}
func UpdateDiagnosis(diagnosisID string, updatedDiagnosis models.Diagnosis) (*mongo.UpdateResult, error) {
	diagnosisCollection := config.DB.Collection("diagnosis")
	objID, err := primitive.ObjectIDFromHex(diagnosisID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedDiagnosis}

	result, err := diagnosisCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating diagnosis:", err)
		return nil, err
	}
	return result, nil
}

func DeleteDiagnosis(diagnosisID string) (*mongo.DeleteResult, error) {
	diagnosisCollection := config.DB.Collection("diagnosis")
	objID, err := primitive.ObjectIDFromHex(diagnosisID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := diagnosisCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting diagnosis:", err)
		return nil, err
	}
	return result, nil
}

func GetAllDiagnosiss() ([]models.Diagnosis, error) {
	diagnosisCollection := config.DB.Collection("diagnosis")

	cursor, err := diagnosisCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting diagnosiss:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var diagnosiss []models.Diagnosis
	for cursor.Next(context.Background()) {
		var diagnosis models.Diagnosis
		if err := cursor.Decode(&diagnosis); err != nil {
			log.Println("Error decoding diagnosis:", err)
			return nil, err
		}
		diagnosiss = append(diagnosiss, diagnosis)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return diagnosiss, nil
}
