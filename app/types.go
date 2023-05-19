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
	breed  string  `json:"breed"`
	Age    int     `json:"age"`
	// Odds float64 `json:"odds"`
}

type Person struct {
	UserName string `json:"user" bson:"user"`
	Pass     string `json:"pass" bson:"pass"`

	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname"  bson:"lasttname"`
	Age       int    `json:"age"  bson:"age"`
	// DOB       string `json:"dob"`
	// Address   string `json:"address"`
}

type Bet struct {
	ID       string  `json:"id"       bson:"_id"`    //int?
	Amount   uint    `json:"amount"   bson:"amount"` //stake
	Profit   int     `json:"profit"   bson:"profit"`
	Investor *Person `json:"investor" bson:"investor"`
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
