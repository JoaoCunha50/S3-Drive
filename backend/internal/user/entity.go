package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Username   string             `bson:"username" json:"username"`
	Role       Role               `bson:"role" json:"role"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	ProfilePic *string            `bson:"profilePic,omitempty" json:"profilePic,omitempty"`
}
