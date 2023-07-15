package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tea struct {
	Type   string
	Rating int32
	Vendor []string `bson:"vendor,omitempty" json:"vendor,omitempty"`
}

type Car struct {
	Id        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	Brand     string             `bson:"brand"`
	Model     string             `bson:"model"`
	Year      int                `bson:"year"`
}

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}
