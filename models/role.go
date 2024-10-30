package models

type Role struct {
    Name        string   `bson:"name" json:"name"`
    Permissions []string `bson:"permissions" json:"permissions"`
}
