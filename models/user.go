package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    Email           string             `bson:"email"`
    Name            string             `bson:"name"`
    Phone           string             `bson:"phone"`
    Address         string             `bson:"address"`
    Gender          string             `bson:"gender"`
    DNI             string             `bson:"dni"`
    Password        string             `bson:"password"`
    BirthDate       string             `bson:"birthDate"`
    Role            UserRole           `bson:"role"`
    Specialist      UserSpecialist     `bson:"specialist"`
    HasAccess       bool               `bson:"hasAccess"`

    
}

type UserRole struct {
    Name        string   `bson:"name"`
    Permissions []string `bson:"permissions"`
}

type UserSpecialist struct {
    Specialization     string `bson:"specialization"`
    Description        string `bson:"description"`
    YearsOfExperience  int    `bson:"yearsOfExperience"`
}
