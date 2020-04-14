package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      string             `json:"user" bson:"user_id"`
	CameraID    string             `json:"camera" bson:"camera_id"`
	Photo       string             `json:"photo" bson:"photo_name"`
	LastUpdated time.Time          `json:"last_updated" bson:"last_updated"`
}
