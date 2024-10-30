// repository/role_repository.go
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

func CreateRole(role models.Role) (*mongo.InsertOneResult, error) {
	roleCollection := config.DB.Collection("role")
	result, err := roleCollection.InsertOne(context.Background(), role)
	if err != nil {
		log.Println("Error creating role:", err)
		return nil, err
	}
	return result, nil
}

func GetRoleByID(roleID string) (*models.Role, error) {
    roleCollection := config.DB.Collection("role") 

    objID, err := primitive.ObjectIDFromHex(roleID)
    if err != nil {
        return nil, err
    }

    var role models.Role
    err = roleCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&role)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        log.Println("Error retrieving role:", err)
        return nil, err
    }
    
    return &role, nil
}
func UpdateRole(roleID string, updatedRole models.Role) (*mongo.UpdateResult, error) {
	roleCollection := config.DB.Collection("role")
	objID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedRole}

	result, err := roleCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating role:", err)
		return nil, err
	}
	return result, nil
}

func DeleteRole(roleID string) (*mongo.DeleteResult, error) {
	roleCollection := config.DB.Collection("role")
	objID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	result, err := roleCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting role:", err)
		return nil, err
	}
	return result, nil
}

func GetAllRoles() ([]models.Role, error) {
	roleCollection := config.DB.Collection("role")

	cursor, err := roleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error getting roles:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var roles []models.Role
	for cursor.Next(context.Background()) {
		var role models.Role
		if err := cursor.Decode(&role); err != nil {
			log.Println("Error decoding role:", err)
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return roles, nil
}
