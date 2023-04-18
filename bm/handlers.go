package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"src/main.go/helpers"

	"src/main.go/types"

	"github.com/gorilla/mux"
)

var (
	horses []types.Horse
	bets   []types.Bet

	// serverPort = 5000
	// mongoPort  = 27017
)

// welcom - html
func Welcom(rw http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("static/index.html"))
	dt := time.Now()

	rw.Header().Set("Content-Type", "text/html")
	tpl.Execute(rw, nil)

	fmt.Fprintf(rw, "<h3 style='color: maroon'>horses available: %v</h3>", len(horses))
	fmt.Fprintf(rw, "<h5 style='color: black'> %s</h5>", dt.Format("01-02-2006 15:04:05 Mon"))
}

// health - html
func Health(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h3 style='color: steelblue'>Got Health</h3>"))
}

// for monitor "OK" loop
func Ping(p int) string {
	healthURL := fmt.Sprintf("http://localhost:%d/health", p)
	resp, _ := http.Get(healthURL)
	SC := resp.StatusCode

	if resp != nil && SC == 200 {
		ret := fmt.Sprintf("OK, Got Health - status code => %d", SC)
		return ret
	} else {
		ret := fmt.Sprintf("Not Ok - status code => %d, (%s)", SC, http.StatusText(SC))
		return ret
	}
}

// get metrics for Monitor func
func getMetrics() string {
	// get the metrics
	metrics := "Horses currently: " + strconv.Itoa(len(horses)) + "\n"
	metrics += "Bets corrently: " + strconv.Itoa(len(bets)) + "\n"
	return metrics
}

func Monitor(rw http.ResponseWriter, r *http.Request) {
	// get the metrics
	metrics := getMetrics()
	// return the metrics
	fmt.Fprintf(rw, "%v", metrics)

	return
}

// func Logger(rw http.ResponseWriter, r *http.Request) {
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
func GetHorses(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	hLen := len(horses)
	json.NewEncoder(rw).Encode(hLen)

	// return the list of horses
	fmt.Fprintf(rw, "%v", horses)
	return
}

// fetch horse by name
func GetHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	horseName := params["name"]

	// check if the horse exists
	horseExists := false
	for _, h := range horses {
		if h.Name == horseName {
			horseExists = true
			break
		}
	}

	if horseExists == false {
		fmt.Fprintf(rw, "Horse named %s does not exist", horseName)
		return
	}

	for _, h := range horses {
		if h.Name == horseName {
			// var horse Horse
			json.NewEncoder(rw).Encode(h)
			return
		}
	}
	json.NewEncoder(rw).Encode(&types.Horse{})

	// // return the horse
	// fmt.Fprintf(rw, "%v", horseName)
	return
}

// Update horse name
func updateHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	horseName := params["name"]
	horseColor := params["color"]

	// duplicate - make into a func
	// check if the horse exists
	horseExists := false
	for _, h := range horses {
		if h.Name == horseName {
			horseExists = true
			break
		}
	}

	if horseExists == false {
		fmt.Fprintf(rw, "Horse named %s does not exist", horseName)
		return
	}

	// newHorse := types.Horse{color = horseColor, name = horseName}

	newHorse := types.Horse{}
	// get the body of the request
	// body, err := ioutil.ReadAll(r.Body)
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		err := json.NewDecoder(r.Body).Decode(&newHorse)
		if err != nil {
			fmt.Fprintf(rw, "Error decoding / reading body")
			log.Printf("err decoding json body. Err: %s", err)
			return
		}

		// check if the new horse info is valid
		if newHorse.Name == "" {
			fmt.Fprintf(rw, "Horse name cannot be empty")
			return
		}
		if newHorse.Color == "" {
			fmt.Fprintf(rw, "Horse color cannot be empty")
			return
		}

		// TODO here probably insert new horse to db
		// update the horse info
		for i, horse := range horses {
			if horse.Name == horseName {
				horses[i].Name = horseName   //newHorse.Name
				horses[i].Color = horseColor //newHorse.Color
				break
			}
		}
		// return the updated horse
		fmt.Fprintf(rw, "%v", horses)
		return
	}

	// for i, item := range horses {
	// 	if item.Name == horseName {
	// 		horses = append(horses[:i], horses[i+1:]...) //redundent
	// 		var horse types.Horse
	// 		_ = json.NewDecoder(r.Body).Decode(&horse)
	// 		horse.Name = horseName
	// 		horses = append(horses, horse)
	// 		json.NewEncoder(rw).Encode(horse)
	// 		return
	// 	}
	// }
}

// bet on a horse
func Invest(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// get the params from the url
	params := mux.Vars(r)
	candidate := params["horse"]
	better := params["investor"]
	// investment, convertionErr := strconv.Atoi(params["amount"])
	investment, convertionErr := params["amount"]
	ID := strconv.Itoa(rand.Intn(100))
	helpers.HandleErr(convertionErr, "Error during str conversion to int.")

	// check if the horse exists
	horseExists := false
	for _, h := range horses {
		if h.Name == candidate {
			horseExists = true
			break
		}
	}

	if horseExists == false {
		fmt.Fprintf(rw, "Horse does not exist")
		return
	}

	// check if the amount is a number or strconv.Atoi() the param (commented above)
	amountIsNumber := false
	amountInt, err := fmt.Sscanf(investment, "%d")
	if err == nil {
		amountIsNumber = true
		if amountInt < 0 {
			fmt.Fprintf(rw, "Amount must be a positive number")
			return
		}
		if amountInt > candidate.Record.Wins {
			fmt.Fprintf(rw, "Amount must be less than horses wins record")
			return
		}
	}

	if amountIsNumber == false {
		fmt.Fprintf(rw, "Amount isnt a number, correct your request")
		return
	}

	// TODO bet check should rely on id checking but now id is randomized.
	// // check if bet is already placed
	// betExists := false
	// for _, bet := range Bets {
	// 	if bet.ID == betID {
	// 		betExists = true
	// 		break
	// 	}
	// }
	// if betExists == true {
	// 	fmt.Fprintf(w, "Bet already placed")
	// 	return
	// }

	for _, h := range horses {
		if h.Name == candidate {
			quotient := h.Record.Loses / h.Record.Wins // https://www.cuemath.com/numbers/quotient/
			// odds := math.Round(quotient)

			// bet := types.Bet{Horse: horseName, Amount: amountInt}
			var bet types.Bet
			bet.ID = ID
			bet.Amount = investment
			bet.Profit = investment * quotient
			bet.Investor.FirstName = better.FirstName // "dvir"
			bet.Investor.LastName = better.LastName   //"gross"

			bets = append(bets, bet) //send to db

			// // update the wins and loses of the horse
			// //win
			// horses[i].Record.Wins += 1 //amountInt
			// //lose
			// horses[i].Record.Loses += 1 //amountInt

			// return the complet new bet
			json.NewEncoder(rw).Encode(bet)
		}
		// return the updated record
		fmt.Fprintf(rw, "%v", horses)

		return
	}
}

func execute() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	}

	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

// func useCPU is a long loop for wasting cpu
func useCPU(w http.ResponseWriter, r *http.Request) {
	count := 1

	for i := 1; i <= 1000000; i++ {
		count = i
	}

	fmt.Printf("count: %d", count)
	w.Write([]byte(fmt.Sprint(count)))
}

func userHandler(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(rw, "The id query parameter is missing", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "<h1>The user id is: %s</h1>", id)
}

func searchHandler(rw http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")

	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)
}

func serve(w http.ResponseWriter, r *http.Request) {
	env := map[string]string{}
	for _, keyval := range os.Environ() {
		keyval := strings.SplitN(keyval, "=", 2)
		env[keyval[0]] = keyval[1]
	}
	bytes, err := json.Marshal(env)
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	w.Write([]byte(bytes))
}

// ###
// ////////
// ###

// // a func to update a horse
// func updateHorse(w http.ResponseWriter, r *http.Request) {

// 	// // check if the horse exists
// 	// horseExists := false
// 	// for _, horse := range horses {
// 	// 	if horse.Name == horseName {
// 	// 		horseExists = true
// 	// 		break

// 	// 	}
// 	// }
// 	// if horseExists == false {
// 	// 	fmt.Fprintf(w, "Horse does not exist")
// 	// 	return
// 	// }

// }

// a func to get the static assets
func staticHandler() http.Handler {
	return http.FileServer(http.Dir("./static"))
}
