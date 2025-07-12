package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name          string             `bson:"name" json:"name"`
	SeName        string             `bson:"seName" json:"seName"`
	ProductsCount int                `bson:"productsCount" json:"productsCount"`
}
