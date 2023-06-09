package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string            `json:"password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	User_Cart       []Product          `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Address struct {
	Address_ID   primitive.ObjectID `json:"_id" bson:"_id"`
	House_Number *string            `json:"house_number"`
	Street       *string            `json:"street"`
	City         *string            `json:"city"`
	Pin_Code     *string            `json:"pin_code"`
}
