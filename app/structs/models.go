package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Horses    map[string]Horse
	MainBoard Board
	Conn      Conn
	ConnErr   error
	Start     time.Time
}

type Record struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Draws  int `json:"draws"`
	Total  int `json:"total"`
}

type Horse struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	Record Record `json:"record"`
	Breed  string `json:"breed"`
	Age    int    `json:"age"`
	// Odds float64 `json:"odds"`
}

type Person struct {
	UserName string `json:"user" bson:"user"`
	Pass     []byte `json:"pass" bson:"pass"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name"  bson:"last_name"`
	Age       int    `json:"age"  bson:"age"`
	// DOB       string `json:"dob"`
	// Address   string `json:"address"`
}

type Bet struct {
	ID         primitive.ObjectID `json:"id"       bson:"_id,omitempty"`
	Amount     uint               `json:"amount"   bson:"amount"` //stake
	Profit     int                `json:"profit"   bson:"profit"`
	InvestorID string             `json:"investor_id" bson:"investor_id"`
	HorseID    int64
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
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
// 	Bets *mongo.Collection `json:"bets" bson:"bets"`
// }

// type Search struct {
// 	Query      string        `json:"query" bson:"query"`
// 	NextPage   int           `json:"next_page" bson:"next_page"`
// 	TotalPages int           `json:"total_pages" bson:"total_pages"`
// 	Results    *news.Results `json:"results" bson:"results"`
// }
