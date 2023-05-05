package main

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Record struct {
	Wins  int `json:"wins"`
	Loses int `json:"loses"`
}

type Horse struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Record *Record `json:"record"`
	// Odds float64 `json:"odds"`
}

type Person struct {
	UserName string `json:"user" bson:"user"`
	Pass     string `json:"pass"`

	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname"  bson:"lasttname"`
	Age       int    `json:"age"  bson:"age"`
	// DOB       string `json:"dob"`
	// Address   string `json:"address"`
}

type Bet struct {
	ID       string  `json:"id"       bson:"_id"`
	Amount   uint    `json:"amount"   bson:"amount"`
	Profit   int     `json:"profit"   bson:"profit"`
	Investor *Person `json:"investor" bson:"investor"`
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
