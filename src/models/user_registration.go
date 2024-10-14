package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct defines the structure of a user in MongoDB.
type User struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name      string             `json:"name" bson:"name"`
    Email     string             `json:"email" bson:"email"`
    Password  string             `json:"password" bson:"password"`
    CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
    UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// Request and response structs
type CreateUserRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email" binding:"omitempty,email"`
}

type UserResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
