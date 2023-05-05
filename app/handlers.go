package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	// "gopkg.in/mgo.v2/bson"

	"rsc.io/quote"
)

// welcom - html
func Welcom(rw http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index2.html"))
	dt := time.Now()
	// user :=

	goQoute := quote.Go()
	// message := fmt.Sprintf(randomFormat(), user)

	rw.Header().Set("Content-Type", "text/html")
	// rw.Header().Set("Accept-Charset","utf-8")
	// rw.Header().Set("Content-Encoding", "gzip")

	tpl.Execute(rw, nil)

	fmt.Fprintf(rw, "<h3 style='color: black'>horses available: %v</h3>", len(horses))
	fmt.Fprintf(rw, "<h3 style='color: black'> random go quote: %v</h3>", goQoute)
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

	if resp != nil && resp.StatusCode == 200 {
		SC := resp.StatusCode
		ret := fmt.Sprintf("OK, Got Health - status code => %d", SC)
		return ret
	} else {
		SC := resp.StatusCode
		ret := fmt.Sprintf("Not Ok - status code => %d, (%s)", SC, http.StatusText(SC))
		return ret
	}
}

// get metrics for Monitor func
func GetMetrics() string {
	// get the metrics
	metrics := "Horses currently: " + strconv.Itoa(len(horses)) + "\n"
	metrics += "Bets corrently: " + strconv.Itoa(len(MainBoard.Bets)) + "\n"
	return metrics
}

func Monitor(rw http.ResponseWriter, r *http.Request) {
	// get the metrics
	metrics := GetMetrics()
	// return the metrics
	fmt.Fprintf(rw, "%v", metrics)

	// return
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

	// return the list of horses
	json.NewEncoder(rw).Encode(horses)

	// fmt.Fprintf(rw, "%v", horses)

	// hLen := len(horses)
	// json.NewEncoder(rw).Encode(hLen)
}

func CreateHorse(w http.ResponseWriter, r *http.Request) {
	// Add your horse creation logic here
}

func DeleteHorse(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	for i, horse := range horses {
		if horse.Name == name {
			horses = append(horses[:i], horses[i+1:]...)
			break
		}
	}

	rw.WriteHeader(http.StatusNoContent)
}

// // fetch horse by name
// func GetHorse(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	horseName := params["name"]

// 	// check if the horse exists
// 	horseExists := false
// 	for _, h := range horses {
// 		if h.Name == horseName {
// 			horseExists = true
// 			break
// 		}
// 	}

// 	if !horseExists {
// 		fmt.Fprintf(rw, "Horse named %s does not exist", horseName)
// 		return
// 	}

// 	for _, h := range horses {
// 		if h.Name == horseName {
// 			// var horse Horse
// 			json.NewEncoder(rw).Encode(h)
// 			return
// 		}
// 	}
// 	json.NewEncoder(rw).Encode(&Horse{})

// 	// // return the horse
// 	// fmt.Fprintf(rw, "%v", horseName)
// 	// return
// }

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

	if !horseExists {
		fmt.Fprintf(rw, "Horse named %s does not exist", horseName)
		return
	}

	// newHorse := Horse{color = horseColor, name = horseName}

	newHorse := Horse{}
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
	// 		var horse Horse
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

	var (
		candidate Horse
		better    Person

		investment int

		params = mux.Vars(r)
	)

	// ***
	// get the params from the url
	ID := strconv.Itoa(rand.Intn(100))
	// id, _ := strconv.Atoi(params["id"])

	candidate.Name = params["horse"]
	betterName := params["investor"]
	better.UserName = betterName
	// better.FirstName = betterName[:]
	// better.FirstName = betterName[:]

	investment, convertionErr := strconv.Atoi(params["amount"])
	// investment, _ = strconv.Atoi(params["amount"])
	Check(convertionErr, "Error during str conversion to int.")

	// ***
	// check if the horse exists
	horseExists := false
	for _, h := range horses {
		if h.Name == candidate.Name {
			horseExists = true
			break
		}
	}

	if !horseExists {
		msg := fmt.Sprintf("Horse %s does not exist", candidate.Name)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

	// ***
	// check if the amount is a number [or strconv.Atoi() the param (commented above)]
	amountIsNumber := false
	// amountInt, err := fmt.Sscanf(investment, "%d")
	// if err == nil {
	// amountIsNumber = true
	// }
	if investment < 0 {
		msg := fmt.Sprintf("Amount must be a positive number, got: %v", investment)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	} else {
		// investment is bigger than 0, valid
		amountIsNumber = true

		//
		if investment > (candidate.Record.Wins * 100) {
			msg := fmt.Sprintf("investment is -> %v, candidate record -> L: %v / W: %v", investment, candidate.Record.Loses, candidate.Record.Wins)
			expln := fmt.Sprintf("Amount must be proportional to %v's record (win / lose) odds", candidate.Name)

			fmt.Fprintf(rw, msg, expln)
			LogToFile(msg + expln)
			return
		}
	}

	if !amountIsNumber {
		msg := "Amount isnt a number, correct your request"
		fmt.Fprint(rw, msg)
		LogToFile(msg)
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
		if h.Name == candidate.Name {
			quotient := h.Record.Loses / h.Record.Wins // https://www.cuemath.com/numbers/quotient/
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

			// return the complet new bet
			json.NewEncoder(rw).Encode(thisBet)
			json.NewEncoder(rw).Encode(MainBoard)
		}
		// return the updated record
		fmt.Fprintf(rw, "%v", horses)

	}
}

// func execute() {
// 	if runtime.GOOS == "windows" {
// 		fmt.Println("Can't Execute this on a windows machine")
// 	}

// 	out, err := exec.Command("ls", "-ltr").Output()
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 	}

// 	fmt.Println("Command Successfully Executed")
// 	output := string(out[:])
// 	fmt.Println(output)
// }

// // func useCPU is a long loop for wasting cpu
// func useCPU(w http.ResponseWriter, r *http.Request) {
// 	count := 1

// 	for i := 1; i <= 1000000; i++ {
// 		count = i
// 	}

// 	fmt.Printf("count: %d", count)
// 	w.Write([]byte(fmt.Sprint(count)))
// }

// func userHandler(rw http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	if id == "" {
// 		http.Error(rw, "The id query parameter is missing", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Fprintf(rw, "<h1>The user id is: %s</h1>", id)
// }

// func searchHandler(rw http.ResponseWriter, r *http.Request) {
// 	u, err := url.Parse(r.URL.String())

// 	if err != nil {
// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	params := u.Query()
// 	searchQuery := params.Get("q")
// 	page := params.Get("page")

// 	if page == "" {
// 		page = "1"
// 	}

// 	fmt.Println("Search Query is: ", searchQuery)
// 	fmt.Println("Page is: ", page)
// }

// func serve(w http.ResponseWriter, r *http.Request) {
// 	env := map[string]string{}
// 	for _, keyval := range os.Environ() {
// 		keyval := strings.SplitN(keyval, "=", 2)
// 		env[keyval[0]] = keyval[1]
// 	}
// 	bytes, err := json.Marshal(env)
// 	if err != nil {
// 		w.Write([]byte("{}"))
// 		return
// 	}
// 	w.Write([]byte(bytes))
// }

// // ###
// // ////////
// // ###

// // // a func to update a horse
// // func updateHorse(w http.ResponseWriter, r *http.Request) {

// // 	// // check if the horse exists
// // 	// horseExists := false
// // 	// for _, horse := range horses {
// // 	// 	if horse.Name == horseName {
// // 	// 		horseExists = true
// // 	// 		break

// // 	// 	}
// // 	// }
// // 	// if horseExists == false {
// // 	// 	fmt.Fprintf(w, "Horse does not exist")
// // 	// 	return
// // 	// }

// // }

// // a func to get the static assets
// func staticHandler() http.Handler {
// 	return http.FileServer(http.Dir("./static"))
// }

// // LogRequest -> logs req info
// func LogRequest(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		color.Yellow("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
// 		handler.ServeHTTP(w, r)
// 	})
// }
