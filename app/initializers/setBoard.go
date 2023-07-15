package inits

import (
	"time"

	"github.com/Slvr-one/bookmaker/structs"
)

func SetBoard(mb structs.Board, start time.Time) {
	mb.Title = "welcom to the Garrison; what we have today: "
	mb.Footer = "hope to see tou here again"
	mb.Date = &start

}
