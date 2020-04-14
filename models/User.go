package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	UserID uint
	jwt.StandardClaims
}

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id" `
	UserID     string             `json:"user_id" bson:"user_id"`
	Username   string             `json:"username" bson:"username"`
	Password   string             `json:"-" bson:"password"`
	Name       string             `json:"name" bson:"name" `
	Phone      string             `json:"phone" bson:"phone" `
	PhotoName  string             `json:"photo_name" bson:"photo" `
	Role       string             `json:"role" bson:"role"`
	Status     int32              `json:"status" bson:"status"`
	Token      string             `json:"-"`
	Embeddings [][]float64        `json:"embeddings" bson:"embeddings"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DeletedAt  *time.Time         `json:"deleted_at"`
}

type UserJSON struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"size:100;unique_index"`
	Name      string `json:"name" gorm:"size:255"`
	PhotoName string `json:"photo_name" gorm:"size:100"`
}

type UserRecognition struct {
	UserID        string    `json:"user_id"`
	Name          string    `json:"name"`
	Accuracy      float64   `json:"accuracy"`
	Elapsed       string    `json:"elapsed"`
	Embedding     []float64 `json:"embeddings"`
	PhotoEncoding string    `json:"image"`
	CameraID      string    `json:"camera_id"`
}

type UserVerify struct {
	UserID     int         `json:"user_id"`
	Embeddings [][]float64 `json:"embeddings"`
}

type UserEmbeddings struct {
	UserID     string      `json:"user_id" bson:"user_id"`
	Name       string      `json:"name" bson:"name"`
	Embeddings [][]float64 `json:"embeddings" bson:"embeddings"`
}
