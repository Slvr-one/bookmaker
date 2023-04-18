package types

import "go.mongodb.org/mongo-driver/mongo"

// diclare structs:

type Record struct {
	Wins  int `json:"wins"`
	Loses int `json:"loses"`
}

type Horse struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Record *Record `json:"record"`
}

type Person struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname"  bson:"lasttname"`
}

type Bet struct {
	ID       string  `json:"id"       bson:"_id"`
	Amount   uint    `json:"amount"   bson:"amount"`
	Profit   int     `json:"profit"   bson:"profit"`
	Investor *Person `json:"investor" bson:"investor"`
}

type DatabaseCollections struct {
	Bets *mongo.Collection
}
