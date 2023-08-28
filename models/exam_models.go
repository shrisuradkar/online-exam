package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID         primitive.ObjectID `bson:"_id"`
	Last_name  *string            `json:"last_name" validate:"required,min=1,max=100"`
	Password   *string            `json:"Password" validate:"required,min=6"`
	Email      *string            `json:"email" validate:"email,required"`
	Phone      *string            `json:"phone" validate:"required"`
	Course     *string            `json:"course" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Exam_id    string             `json:"exam_id"`
}
