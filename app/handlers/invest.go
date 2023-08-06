package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	s "github.com/Slvr-one/bookmaker/app/structs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// bet on a horse
func Invest(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	var (
		candidate  s.Horse
		better     s.Person
		investment int
		params     = mux.Vars(ctx.Request)
	)

	// ***
	// get the params from the url
	ID := strconv.Itoa(rand.Intn(100))
	// id, _ := strconv.Atoi(params["id"])

	// par := fmt.Sprintf("params: %v \n", params)
	// fmt.Fprint(ctx.Writer, par)

	betterName := params["investor"]
	candidate.Name = params["horse"]
	better.UserName = betterName
	// better.FirstName = betterName[:]
	// better.FirstName = betterName[:]

	investment, convertionErr := strconv.Atoi(params["amount"])
	// investment, _ = strconv.Atoi(params["amount"])
	msg := fmt.Sprintf("Error conversion str %v to int", params["amount"])
	Check(convertionErr, msg)

	// ***
	// check if the horse exists
	exist := IfHorseExist(candidate.Name, Horses)

	if !exist {
		msg := fmt.Sprintf("Horse named %s does not exist", candidate.Name)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	// ***
	amountPositive := false

	// TODO bet check should rely on id checking but now id is randomized.
	// check if the amount is a number [or strconv.Atoi() the param (commented above)]
	// amountIsNumber := false
	// amountInt, err := fmt.Sscanf(investment, "%d")
	// if err == nil {
	// amountIsNumber = true
	// }
	if investment < 0 {
		msg := fmt.Sprintf("Amount must be a positive number, got: %v", investment)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	} else {
		// investment is bigger than 0, valid
		// amountIsNumber = true
		amountPositive = true

		//
		if investment > (candidate.Record.Wins * 100) {
			msg := fmt.Sprintf("investment is -> %v, candidate record -> L: %v / W: %v", investment, candidate.Record.Losses, candidate.Record.Wins)
			expln := fmt.Sprintf("Amount must be proportional to %v's record (win / lose) odds", candidate.Name)

			fmt.Fprintf(ctx.Writer, msg, expln)
			LogToFile(msg + expln)
			return
		}
	}

	if !amountPositive {
		msg := "Amount isnt a positive number, correct your request"
		// msg := "Amount isnt a number, correct your request"
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	for _, h := range Horses {
		if h.Name == candidate.Name {
			quotient := h.Record.Losses / h.Record.Wins // https://www.cuemath.com/numbers/quotient/
			// odds := math.Round(quotient)
			// var con string = "hsfj"

			currentBets := len(MainBoard.Bets)
			thisBet := MainBoard.Bets[currentBets]

			// bet := Bet{Horse: horseName, Amount: amountInt}
			// var bet Bet
			thisBet.ID = ID
			thisBet.Amount = uint(investment)
			thisBet.Profit = investment * quotient
			thisBet.Investor.FirstName = better.FirstName // "dvir"
			thisBet.Investor.LastName = better.LastName   //"gross"

			// MainBoard.Bets = append(MainBoard.Bets, bet) //send to db

			// // update the wins and loses of the horse
			// //win
			// horses[i].Record.Wins += 1 //amountInt
			// //lose
			// horses[i].Record.Loses += 1 //amountInt

			result, insertErr := Conn.Collection.InsertOne(ctx.Request.Context(), thisBet)
			Check(insertErr, "err on db objects insertion.")

			msg := fmt.Sprintf("Inserted a single document: %v", result)
			LogToFile(msg)
			// return the complet new bet
			json.NewEncoder(ctx.Writer).Encode(thisBet)
			json.NewEncoder(ctx.Writer).Encode(MainBoard)
		}
		// return the updated record
		fmt.Fprintf(ctx.Writer, "%v", Horses)
		fmt.Fprintf(ctx.Writer, "%v", MainBoard)
	}
}
