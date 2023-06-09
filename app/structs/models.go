package main

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Record struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Draws  int `json:"draws"`
	Total  int `json:"total"`
}

type Horse struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Record *Record `json:"record"`
	Breed  string  `json:"breed"`
	Age    int     `json:"age"`
	// Odds float64 `json:"odds"`
}

type Person struct {
	UserName string `json:"user" bson:"user"`
	Pass     string `json:"pass" bson:"pass"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name"  bson:"last_name"`
	Age       int    `json:"age"  bson:"age"`
	// DOB       string `json:"dob"`
	// Address   string `json:"address"`
}

type Bet struct {
	// ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	ID        string  `json:"id"       bson:"_id"`    //int?
	Amount    uint    `json:"amount"   bson:"amount"` //stake
	Profit    int     `json:"profit"   bson:"profit"`
	Investor  *Person `json:"investor" bson:"investor"`
	HorseID   int64
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	// Date     time.Time `json:"date" bson:"date"`
	// Status   string  `json:"status" bson:"status"`
}

type Board struct {
	Title  string
	Date   *time.Time
	Bets   [100]Bet
	Footer string
}

type Conn struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// type DatabaseCollections struct {
// 	Bets *mongo.Collection
// }

// type Search struct {
// 	Query      string
// 	NextPage   int
// 	TotalPages int
// 	Results    *news.Results
// }
