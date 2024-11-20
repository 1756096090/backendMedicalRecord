package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    Email           string             `bson:"email"`
    Name            string             `bson:"name"`
    Phone           string             `bson:"phone"`
    Address         string             `bson:"address"`
    Gender          string             `bson:"gender"`
    DNI             string             `bson:"dni"`
    Password        string             `bson:"password"`
    BirthDate       string             `bson:"birthDate"`
    RoleID          primitive.ObjectID `bson:"roleId"` 
    SpecialistID    primitive.ObjectID `bson:"specialistId"`
    HasAccess       bool               `bson:"hasAccess"`
    Role            Role               `bson:"role"`           
    Specialist      Specialist         `bson:"specialist"`     
}