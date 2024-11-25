// repository/schedule_repository.go
package repository

import (
	"backendMedicalRecord/config"
	"backendMedicalRecord/models"
	"context"
	"log"
	"time"
	"fmt"
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


func GetShedulesByMonthYear(month int, year int) ([]models.Schedule, error) {
	scheduleCollection := config.DB.Collection("schedule")
	// time.Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)
	

	filter := bson.M{
		"start_appointment": bson.M{
			"$gte": startDate, 
			"$lt":  endDate,
		},
	}
	cursor, err:= scheduleCollection.Find(context.Background(), filter)
	if err != nil{
		log.Println("Error getting schedules by month and year:", err)
        return nil, err
	}
	defer cursor.Close(context.Background())

	var schedules []models.Schedule
	for cursor.Next(context.Background()){
		var schedule models.Schedule
        if err := cursor.Decode(&schedule); err!= nil{
            log.Println("Error decoding schedule:", err)
            return nil, err
        }
        schedules = append(schedules, schedule)
	}

	if err := cursor.Err(); err!= nil{
        log.Println("Cursor error:", err)
        return nil, err
    }

	return schedules, nil
}


func GetSchedulesByUserAndDate(userID string) ([]models.ScheduleDetail, error) {
	scheduleCollection := config.DB.Collection("schedule")

	// Calculate date range
	now := time.Now().UTC()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"id_user", "673e5b3670412d1df3072a46"},
				{"start_appointment", bson.M{
					"$gte": startDate, // Filter for appointments after startDate
					"$lt":  endDate,   // Filter for appointments before endDate
				}},
				

			}},
		},
		{
			{"$addFields", bson.D{
				{"id_user", bson.D{{"$toObjectId", "$id_user"}}},
				{"id_patient", bson.D{{"$toObjectId", "$id_patient"}}},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", "user"},
				{"localField", "id_user"},
				{"foreignField", "_id"},
				{"as", "user"},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", "patient"},
				{"localField", "id_patient"},
				{"foreignField", "_id"},
				{"as", "patient"},
			}},
		},
		{
			{"$unwind", "$user"},
		},
		{
			{"$unwind", "$patient"},
		},
	}
	
	

	// Execute aggregation
	cursor, err := scheduleCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Printf("Error executing aggregation: %v", err)
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.Background())

	// Decode results
	var schedules []models.ScheduleDetail
	if err := cursor.All(context.Background(), &schedules); err != nil {
		log.Printf("Error decoding schedules: %v", err)
		return nil, fmt.Errorf("failed to decode schedules: %w", err)
	}

	return schedules, nil
}