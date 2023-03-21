package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	Order_ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Order_Cart     []Product          `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at"`
	Price          int                `json:"price"`
	Discount       *int               `json:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
