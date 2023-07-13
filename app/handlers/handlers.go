package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	s "github.com/Slvr-one/bookmaker/structs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"rsc.io/quote"
)

// welcom - html
func Welcom(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl := template.Must(template.ParseFiles("templates/index2.html"))
	dt := time.Now()
	// user :=

	goQoute := quote.Go()
	// message := fmt.Sprintf(randomFormat(), user)

	tpl.Execute(ctx.Writer, nil)
	// io.WriteString(ctx.Writer, "Hello again, Gefyra!")

	var horses []s.Horse //handle with func / ctx
	fmt.Fprintf(ctx.Writer, "<h3 style='color: black'>horses available: %v</h3>", horses)
	fmt.Fprintf(ctx.Writer, "<h3 style='color: black'> random go quote: %v</h3>", goQoute)
	fmt.Fprintf(ctx.Writer, "<h5 style='color: black'> %s</h5>", dt.Format("01-02-2006 15:04:05 Mon"))
}

// health - html
func Health(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/html")
	ctx.Writer.Write([]byte("<h3 style='color: steelblue'>Got Health</h3>"))

}

// for monitor "OK" loop
func Ping(host string, port int) string {
	// if p == 0 {
	//     return "OK"
	// }
	// return "ERROR"

	healthURL := fmt.Sprintf("http://%v:%d/health", host, port)

	resp, getErr := http.Get(healthURL)
	msg := fmt.Sprintf("err on http get for pinging %v", healthURL)
	Check(getErr, msg)
	if resp != nil {
		defer resp.Body.Close()
		SC := resp.StatusCode

		// check status code:
		if SC == 200 {
			ret := fmt.Sprintf("OK, Got Health - status code => %d, (%s)", SC, http.StatusText(SC))
			return ret
		}
		ret := fmt.Sprintf("Not Ok - status code => %d, (%s)", SC, http.StatusText(SC))
		return ret

	}
	return fmt.Sprintf("ERROR, resp = %v", resp)
}

// get metrics for Monitor func
func GetMetrics(host string, port int) string {
	metricsURL := fmt.Sprintf("http://%v:%d/metrics", host, port)

	resp, getErr := http.Get(metricsURL)
	msg := fmt.Sprintf("err on http get for pinging %v", metricsURL)
	Check(getErr, msg)
	if resp != nil {
		defer resp.Body.Close()
		// SC := resp.StatusCode
	}
	// get the metrics
	metrics := "Horses currently: " + strconv.Itoa(len(Horses)) + "\n"
	metrics += "Bets corrently: " + strconv.Itoa(len(MainBoard.Bets)) + "\n"
	return metrics
}

func Monitor(ctx *gin.Context) {

	// port :=
	host := ctx.Request.Host
	port, _ := strconv.Atoi(DefaultServerPort) //r.URL.Query().Get("port")

	m := GetMetrics(host, port)
	fmt.Fprintf(ctx.Writer, "%v", m)
}

// func Logger(ctx *gin.Context) {
// 	//if app is up, keep logging every 5
// 	for {
// 		// result := Ping()
// 		// fmt.Println(result)
// 		fmt.Printf("\r%s -- %d", "logs", rand.Intn(100000000)) //strconv.Itoa(rand.Intn(100000000)))
// 		time.Sleep(5 * time.Second)
// 	}
// }

///////////////////////////////////////////////////////////

// fetch horses available (number)
func GetHorses(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	// return the list of horses
	json.NewEncoder(ctx.Writer).Encode(Horses)
	// fmt.Fprintf(ctx.Writer, "%v", horses)

	// hLen := len(horses)
	// json.NewEncoder(ctx.Writer).Encode(hLen)
}

func CreateHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	exist := IfHorseExist(horseName, Horses)

	if !exist {
		msg := fmt.Sprintf("Cannot create, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

}

func DeleteHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot delete, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	for i, horse := range horses {
		if horse.Name == horseName {
			/* append() adds elements to a slice, but in this case, its used to remove an element:
			The [:i] notation denotes that all the elements before the i-th index are included,
			while [i+1:] denotes that all the elements after the i-th index are included.
			By combining these two notations with the ellipsis(...), the horses slice is effectively modified to exclude the i-th element.
			note that the original slice is not modified; rather, a new slice is created missing the i-th element. */
			horses = append(horses[:i], horses[i+1:]...)
			break
		}
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

// fetch horse by name
func GetHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	// check if the horse exists
	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot get, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	for _, h := range horses {

		if h.Name == horseName {
			// var horse Horse
			json.NewEncoder(ctx.Writer).Encode(h)
			return
		}
	}
	// json.NewEncoder(ctx.Writer).Encode(&Horse{})

	// // return the horse
	// fmt.Fprintf(ctx.Writer, "%v", horseName)
	// return
}

// Update horse name
func UpdateHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]
	horseColor := params["color"]
	horseAge, _ := strconv.Atoi(params["age"])

	// check if the horse exists
	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot update, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	newHorse := Horse{Name: horseName, Color: horseColor, Age: horseAge}

	// newHorse := Horse{color = horseColor, name = horseName}

	// get the body of the request
	// body, err := ioutil.ReadAll(r.Body)
	_, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		err := json.NewDecoder(ctx.Request.Body).Decode(&newHorse)
		if err != nil {
			msg := fmt.Sprintf("Error decoding / reading json body - %v", err)
			fmt.Fprint(ctx.Writer, msg)
			Check(err, msg)
			return
		}

		// check if the new horse info is valid
		if newHorse.Name == "" {
			fmt.Fprintf(ctx.Writer, "Horse name cannot be empty")
			return
		}
		if newHorse.Color == "" {
			fmt.Fprintf(ctx.Writer, "Horse color cannot be empty")
			return
		}

		// TODO here probably insert new horse to db
		// update the horse info

		for i, h := range horses {
			if h.Name == horseName {
				horses[i].Name = horseName   //newHorse.Name
				horses[i].Color = horseColor //newHorse.Color
				horses[i].Age = horseAge     //newHorse.Age
				horses[i].Record.Losses = 0
				horses[i].Record.Wins = 0

				// horses[i].Record = &Record{0, 0}

				break
			}
		}
		// return the updated horse

		fmt.Fprintf(ctx.Writer, "%v", horses)
		return
	}

	// for i, item := range horses {
	// 	if item.Name == horseName {
	// 		horses = append(horses[:i], horses[i+1:]...) //redundent
	// 		var horse Horse
	// 		_ = json.NewDecoder(r.Body).Decode(&horse)
	// 		horse.Name = horseName
	// 		horses = append(horses, horse)
	// 		json.NewEncoder(ctx.Writer).Encode(horse)
	// 		return
	// 	}
	// }
}

// bet on a horse
func Invest(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	var (
		candidate  Horse
		better     Person
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
	exist := ifHorseExist(candidate.Name)

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

	for _, h := range horses {
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

			result, insertErr := conn.Collection.InsertOne(ctx.Request.Context(), thisBet)
			Check(insertErr, "err on db objects insertion.")

			msg := fmt.Sprintf("Inserted a single document: %v", result)
			LogToFile(msg)
			// return the complet new bet
			json.NewEncoder(ctx.Writer).Encode(thisBet)
			json.NewEncoder(ctx.Writer).Encode(MainBoard)
		}
		// return the updated record
		fmt.Fprintf(ctx.Writer, "%v", horses)
		fmt.Fprintf(ctx.Writer, "%v", MainBoard)
	}
}
