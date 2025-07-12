package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vendor struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `bson:"name" json:"name"`
	SeName         string             `bson:"seName" json:"seName"`
	ImageUrl       string             `bson:"imageUrl" json:"imageUrl"`
	ProductCount   int                `bson:"productCount" json:"ProductCount"`
	FollowersCount int                `bson:"followersCount" json:"followersCount"`
}
