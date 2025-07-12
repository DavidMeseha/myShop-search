package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name         string             `bson:"name" json:"name"`
	SeName       string             `bson:"seName" json:"seName"`
	ProductCount int                `bson:"productCount" json:"productCount"`
}
