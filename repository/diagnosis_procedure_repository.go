// repository/diagnosisProcedure_repository.go
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

func CreateDiagnosisProcedure(diagnosisProcedure models.DiagnosisProcedure) (*mongo.InsertOneResult, error) {
	diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure")
	result, err := diagnosisProcedureCollection.InsertOne(context.Background(), diagnosisProcedure)
	if err != nil {
		log.Println("Error creating diagnosisProcedure:", err)
		return nil, err
	}
	return result, nil
}

func GetDiagnosisProcedureByID(diagnosisProcedureID string) (*models.DiagnosisProcedure, error) {
    diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure") 

    objID, err := primitive.ObjectIDFromHex(diagnosisProcedureID)
    if err != nil {
        return nil, err
    }

    var diagnosisProcedure models.DiagnosisProcedure
    err = diagnosisProcedureCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&diagnosisProcedure)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving diagnosisProcedure:", err)
        return nil, err
    }
    
    return &diagnosisProcedure, nil
}
func UpdateDiagnosisProcedure(diagnosisProcedureID string, updatedDiagnosisProcedure models.DiagnosisProcedure) (*mongo.UpdateResult, error) {
	diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure")
	objID, err := primitive.ObjectIDFromHex(diagnosisProcedureID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedDiagnosisProcedure}

	result, err := diagnosisProcedureCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating diagnosisProcedure:", err)
		return nil, err
	}
	return result, nil
}

func DeleteDiagnosisProcedure(diagnosisProcedureID string) (*mongo.DeleteResult, error) {
	diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure")
	objID, err := primitive.ObjectIDFromHex(diagnosisProcedureID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := diagnosisProcedureCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting diagnosisProcedure:", err)
		return nil, err
	}
	return result, nil
}

func GetAllDiagnosisProcedures() ([]models.DiagnosisProcedure, error) {
	diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure")

	cursor, err := diagnosisProcedureCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting diagnosisProcedures:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var diagnosisProcedures []models.DiagnosisProcedure
	for cursor.Next(context.Background()) {
		var diagnosisProcedure models.DiagnosisProcedure
		if err := cursor.Decode(&diagnosisProcedure); err != nil {
			log.Println("Error decoding diagnosisProcedure:", err)
			return nil, err
		}
		diagnosisProcedures = append(diagnosisProcedures, diagnosisProcedure)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return diagnosisProcedures, nil
}


func GetAllDiagnosisProceduresByID(id string) ([]models.DiagnosisProcedure,error){
	diagnosisProcedureCollection := config.DB.Collection("diagnosisProcedure")
	log.Println(id,"get all diagnosis")
    filter := bson.M{"id_patient": id}

    cursor, err := diagnosisProcedureCollection.Find(context.Background(), filter)
    if err!= nil {
        log.Println("Error getting diagnosisProcedures:", err)
        return nil, err
    }
    defer cursor.Close(context.Background())

    var diagnosisProcedures []models.DiagnosisProcedure
    for cursor.Next(context.Background()) {
        var diagnosisProcedure models.DiagnosisProcedure
        if err := cursor.Decode(&diagnosisProcedure); err!= nil {
            log.Println("Error decoding diagnosisProcedure:", err)
            return nil, err
        }
        diagnosisProcedures = append(diagnosisProcedures, diagnosisProcedure)
    }

	if err := cursor.Err(); err!= nil{
        log.Println("Cursor error:", err)
        return nil, err
    }
	return diagnosisProcedures, nil
}