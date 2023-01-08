package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_name      *string            `json:"first_name" validate:"required,min=2,max=30`
	Last_name       *string            `json:"last_name" validate:"required,min=2,max=30`
	Password        *string            `json:"password"  validate:"required,min=6"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_token   *string            `json:"refresh_token"`
	Create_At       time.Time          `json:"created_at"`
	Update_At       time.Time          `json:"updated_at"`
	User_ID         *string            `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart`
	Address_details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"address" bson:"address"`
}
