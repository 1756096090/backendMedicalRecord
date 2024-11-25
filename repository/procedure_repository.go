// repository/procedure_repository.go
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

func CreateProcedure(procedure models.Procedure) (*mongo.InsertOneResult, error) {
	procedureCollection := config.DB.Collection("procedure")
	result, err := procedureCollection.InsertOne(context.Background(), procedure)
	if err != nil {
		log.Println("Error creating procedure:", err)
		return nil, err
	}
	return result, nil
}

func GetProcedureByID(procedureID string) (*models.Procedure, error) {
    procedureCollection := config.DB.Collection("procedure") 

    objID, err := primitive.ObjectIDFromHex(procedureID)
    if err != nil {
        return nil, err
    }

    var procedure models.Procedure
    err = procedureCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&procedure)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving procedure:", err)
        return nil, err
    }
    
    return &procedure, nil
}
func UpdateProcedure(procedureID string, updatedProcedure models.Procedure) (*mongo.UpdateResult, error) {
	procedureCollection := config.DB.Collection("procedure")
	objID, err := primitive.ObjectIDFromHex(procedureID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedProcedure}

	result, err := procedureCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating procedure:", err)
		return nil, err
	}
	return result, nil
}

func DeleteProcedure(procedureID string) (*mongo.DeleteResult, error) {
	procedureCollection := config.DB.Collection("procedure")
	objID, err := primitive.ObjectIDFromHex(procedureID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := procedureCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting procedure:", err)
		return nil, err
	}
	return result, nil
}

func GetAllProcedures() ([]models.Procedure, error) {
	procedureCollection := config.DB.Collection("procedure")

	cursor, err := procedureCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting procedures:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var procedures []models.Procedure
	for cursor.Next(context.Background()) {
		var procedure models.Procedure
		if err := cursor.Decode(&procedure); err != nil {
			log.Println("Error decoding procedure:", err)
			return nil, err
		}
		procedures = append(procedures, procedure)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return procedures, nil
}
