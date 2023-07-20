package db

import (
	"fmt"

	h "github.com/Slvr-one/bookmaker/handlers"
	"github.com/Slvr-one/bookmaker/structs"
	"github.com/gin-gonic/gin"
)

func GetAvHorses(conn structs.Conn, ctx *gin.Context) {
	dbName, collName := "bookmaker", "bets"
	coll := conn.Client.Database(dbName).Collection(collName)
	docs := []interface{}{
		structs.Tea{Type: "Masala", Rating: 10, Vendor: []string{"A", "C"}},
		structs.Tea{Type: "English Breakfast", Rating: 6},
		structs.Tea{Type: "Oolong", Rating: 7, Vendor: []string{"C"}},
		structs.Tea{Type: "Assam", Rating: 5},
		structs.Tea{Type: "Earl Grey", Rating: 8, Vendor: []string{"A", "B"}},
	}
	result, err := coll.InsertMany(ctx, docs)
	h.Check(err, "checking")
	fmt.Println(result)

}
