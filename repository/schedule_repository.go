// repository/schedule_repository.go
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

func CreateSchedule(schedule models.Schedule) (*mongo.InsertOneResult, error) {
	scheduleCollection := config.DB.Collection("schedule")
	result, err := scheduleCollection.InsertOne(context.Background(), schedule)
	if err != nil {
		log.Println("Error creating schedule:", err)
		return nil, err
	}
	return result, nil
}

func GetScheduleByID(scheduleID string) (*models.Schedule, error) {
    scheduleCollection := config.DB.Collection("schedule") 

    objID, err := primitive.ObjectIDFromHex(scheduleID)
    if err != nil {
        return nil, err
    }

    var schedule models.Schedule
    err = scheduleCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&schedule)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving schedule:", err)
        return nil, err
    }
    
    return &schedule, nil
}
func UpdateSchedule(scheduleID string, updatedSchedule models.Schedule) (*mongo.UpdateResult, error) {
	scheduleCollection := config.DB.Collection("schedule")
	objID, err := primitive.ObjectIDFromHex(scheduleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedSchedule}

	result, err := scheduleCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating schedule:", err)
		return nil, err
	}
	return result, nil
}

func DeleteSchedule(scheduleID string) (*mongo.DeleteResult, error) {
	scheduleCollection := config.DB.Collection("schedule")
	objID, err := primitive.ObjectIDFromHex(scheduleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := scheduleCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting schedule:", err)
		return nil, err
	}
	return result, nil
}

func GetAllSchedules() ([]models.Schedule, error) {
	scheduleCollection := config.DB.Collection("schedule")

	cursor, err := scheduleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting schedules:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var schedules []models.Schedule
	for cursor.Next(context.Background()) {
		var schedule models.Schedule
		if err := cursor.Decode(&schedule); err != nil {
			log.Println("Error decoding schedule:", err)
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return schedules, nil
}
