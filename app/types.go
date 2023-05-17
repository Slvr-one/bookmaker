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

// type Game struct {
// 	Players []Player         // A list of all active players in the game
// 	Horses  map[string]Horse // A mapping from each horse name to its details
// 	Stakes  map[int64]*Stake // A mapping from each player ID to their current stake
// }

var Html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="description" content="This is a demo of the Quote API">
<meta name="go-import" content="github.com/rsc/quote"
<style>
body {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 14px;
	font-weight: 400;
	color: #000000;
	background-color: #fff;
	padding: 0;
	margin: 0;
	text-align: center;
	text-decoration: none;
	border: 1px solid #ccc;
	border-radius: 5px;
	border-color: #ccc;
	border-style: solid;
	border-width: 1px;
	border-top-left-radius: 5px;
	border-top-right-radius: 5px;
	border-bottom-left-radius: 5px;
	border-bottom-right-radius: 5px;
	border-top: 1px solid #ccc;
	border-bottom: 1px solid #ccc;
	border-left: 1px solid #ccc;
	border-right: 1px solid #ccc;
}
.container {
	margin: 0;
	padding: 0;
}
.quote {
	margin: 0;
	padding: 0;
}
</style>
</head>
<body>
<div class="container">
<div class="quote">
<h1>Welcome to the Quote API!</h1>
<p>This is a demo of the Quote API. You can find the source code <a href="https://github.com/rsc/quote">here</a>.</p>
</div>
</body>
</html>`
