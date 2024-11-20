// repository/user_repository.go
package repository

import (
	"backendMedicalRecord/config"
	"backendMedicalRecord/models"
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	userCollection := config.DB.Collection("user")
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}
	return result, nil
}

func GetUserByID(userID string) (*models.User, error) {
    userCollection := config.DB.Collection("user") 

    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return nil, err 
    }

    var user models.User
    err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving user:", err)
        return nil, err
    }
    
    return &user, nil
}
func UpdateUser(userID string, updatedUser models.User) (*mongo.UpdateResult, error) {
	userCollection := config.DB.Collection("user")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedUser}

	result, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating user:", err)
		return nil, err
	}
	return result, nil
}


func DeleteUser(userID string) (*mongo.DeleteResult, error) {
	userCollection := config.DB.Collection("user")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting user:", err)
		return nil, err
	}
	return result, nil
}

func GetAllUsers() ([]models.User, error) {
	userCollection := config.DB.Collection("user")

	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting users:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return users, nil
}


func Login(email string, password string) (models.User, error) {
    collection := config.DB.Collection("user")
    
    var user models.User

    err := collection.FindOne(context.Background(), bson.M{"email": email, "password": password}).Decode(&user)
    if err != nil {
        fmt.Println("Error finding user:", err)
        return user, err
    }

    return user, nil
}

func GetUserByDNIOrEmail(dni string, email string) (*models.User, error) {
    userCollection := config.DB.Collection("user")

    var user models.User
    filter := bson.M{"$or": []bson.M{
        {"dni": dni},
        {"email": email},
    }}

    err := userCollection.FindOne(context.Background(), filter).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        log.Println("Error retrieving user by DNI or email:", err)
        return nil, err
    }

    return &user, nil
}




func GetUsersAllInfo() ([]models.UserInfo, error) {
    userCollection := config.DB.Collection("user")

    pipeline := mongo.Pipeline{
        {
            {"$lookup", bson.D{
                {"from", "role"}, 
                {"localField", "roleId"},
                {"foreignField", "_id"},
                {"as", "role"},
            }},
        },
        {
            {Key: "$lookup", Value: bson.D{
                {"from", "specialist"}, 
                {"localField", "specialistId"},
                {"foreignField", "_id"},
                {"as", "specialist"},
            }},
        },
        {
            {"$unwind", "$role"}, 
        },
        {
            {"$unwind", "$specialist"}, 
        },
    }

    cursor, err := userCollection.Aggregate(context.Background(), pipeline)
    if err != nil {
        log.Println("Error getting users:", err)
        return nil, err
    }
    defer cursor.Close(context.Background())

    var users []models.UserInfo
    for cursor.Next(context.Background()) {
        var user models.UserInfo
        if err := cursor.Decode(&user); err != nil {
            log.Println("Error decoding user:", err)
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        log.Println("Cursor error:", err)
        return nil, err
    }

    return users, nil
}


