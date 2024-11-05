// repository/specialist_repository.go
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

func CreateSpecialist(specialist models.Specialist) (*mongo.InsertOneResult, error) {
	specialistCollection := config.DB.Collection("specialist")
	result, err := specialistCollection.InsertOne(context.Background(), specialist)
	if err != nil {
		log.Println("Error creating specialist:", err)
		return nil, err
	}
	return result, nil
}

func GetSpecialistByID(specialistID string) (*models.Specialist, error) {
    specialistCollection := config.DB.Collection("specialist") 

    objID, err := primitive.ObjectIDFromHex(specialistID)
    if err != nil {
        return nil, err
    }

    var specialist models.Specialist
    err = specialistCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&specialist)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving specialist:", err)
        return nil, err
    }
    
    return &specialist, nil
}
func UpdateSpecialist(specialistID string, updatedSpecialist models.Specialist) (*mongo.UpdateResult, error) {
	specialistCollection := config.DB.Collection("specialist")
	objID, err := primitive.ObjectIDFromHex(specialistID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedSpecialist}

	result, err := specialistCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating specialist:", err)
		return nil, err
	}
	return result, nil
}

func DeleteSpecialist(specialistID string) (*mongo.DeleteResult, error) {
	specialistCollection := config.DB.Collection("specialist")
	objID, err := primitive.ObjectIDFromHex(specialistID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := specialistCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting specialist:", err)
		return nil, err
	}
	return result, nil
}

func GetAllSpecialists() ([]models.Specialist, error) {
	specialistCollection := config.DB.Collection("specialist")

	cursor, err := specialistCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting specialists:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var specialists []models.Specialist
	for cursor.Next(context.Background()) {
		var specialist models.Specialist
		if err := cursor.Decode(&specialist); err != nil {
			log.Println("Error decoding specialist:", err)
			return nil, err
		}
		specialists = append(specialists, specialist)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return specialists, nil
}
