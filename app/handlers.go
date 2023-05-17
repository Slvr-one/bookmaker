package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	// "html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	// "time"

	"github.com/gorilla/mux"
	"rsc.io/quote"

	// "gopkg.in/mgo.v2/bson"
	// "rsc.io/quote"
	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// welcom - html
func Welcom(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	// rw.Header().Set("Content-Type", "text/html")
	// rw.Header().Set("Accept-Charset","utf-8")

	// rw.Header().Set("Cache-Control", "no-cache")
	// rw.Header().Set("Cache-Control","no-cache, no-store, must-revalidate")

	// rw.Header().Set("Pragma","no-cache")
	// rw.Header().Set("Expires","0")
	// rw.Header().Set("Access-Control-Allow-Origin","*")
	// rw.Header().Set("Access-Control-Allow-Headers","Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// rw.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
	// rw.Header().Set("Access-Control-Expose-Headers","Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// rw.Header().Set("Access-Control-Allow-Credentials","true")
	// rw.Header().Set("Content-Encoding", "gzip")
	// rw.Header().Set("Connection", "keep-alive")

	// fmt.Fprintf(rw, Html)

	tpl := template.Must(template.ParseFiles("templates/index2.html"))
	dt := time.Now()
	// user :=

	goQoute := quote.Go()
	// message := fmt.Sprintf(randomFormat(), user)

	tpl.Execute(rw, nil)
	// io.WriteString(rw, "Hello again, Gefyra!")

	fmt.Fprintf(rw, "<h3 style='color: black'>horses available: %v</h3>", len(horses))
	fmt.Fprintf(rw, "<h3 style='color: black'> random go quote: %v</h3>", goQoute)
	fmt.Fprintf(rw, "<h5 style='color: black'> %s</h5>", dt.Format("01-02-2006 15:04:05 Mon"))
}

// health - html
func Health(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h3 style='color: steelblue'>Got Health</h3>"))

}

func httpErrorBadRequest(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusBadRequest)
}

func httpErrorInternalServerError(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusInternalServerError)
}

func httpError(err error, ctx *gin.Context, status int) {
	log.Println(err.Error())
	ctx.String(status, err.Error())
}

func pingHandler(ctx *gin.Context) {
	req := resty.New().R().SetHeader("Content-Type", "application/text")
	otelCtx := ctx.Request.Context()
	span := trace.SpanFromContext(otelCtx)
	defer span.End()
	otel.GetTextMapPropagator().Inject(otelCtx, propagation.HeaderCarrier(req.Header))
	url := ctx.Query("url")
	if len(url) == 0 {
		url = os.Getenv("PING_URL")
		if len(url) == 0 {
			httpErrorBadRequest(errors.New("url is empty"), ctx)
			return
		}
	}
	log.Printf("Sending a ping to %s", url)
	resp, err := req.Get(url)
	if err != nil {
		httpErrorBadRequest(err, ctx)
		return
	}
	log.Println(resp.String())
	ctx.String(http.StatusOK, resp.String())
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
	metrics := "Horses currently: " + strconv.Itoa(len(horses)) + "\n"
	metrics += "Bets corrently: " + strconv.Itoa(len(MainBoard.Bets)) + "\n"
	return metrics
}

func Monitor(rw http.ResponseWriter, r *http.Request) {

	// port :=
	host := r.Host
	port, _ := strconv.Atoi(defaultServerPort) //r.URL.Query().Get("port")

	m := GetMetrics(host, port)
	fmt.Fprintf(rw, "%v", m)
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

func CreateHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	horseName := params["name"]

	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot create, Horse named %s does not exist", horseName)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

}

func DeleteHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	horseName := params["name"]

	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot delete, Horse named %s does not exist", horseName)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

	for i, horse := range horses {
		if horse.Name == horseName {
			// append() adds elements to a slice, but in this case, its used to remove an element:
			// The [:i] notation denotes that all the elements before the i-th index are included,
			// while [i+1:] denotes that all the elements after the i-th index are included.
			// By combining these two notations with the ellipsis(...), the horses slice is effectively modified to exclude the i-th element.
			// note that the original slice is not modified; rather, a new slice is created missing the i-th element.
			horses = append(horses[:i], horses[i+1:]...)
			break
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}

// fetch horse by name
func GetHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	horseName := params["name"]

	// check if the horse exists
	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot get, Horse named %s does not exist", horseName)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

	for _, h := range horses {

		if h.Name == horseName {
			// var horse Horse
			json.NewEncoder(rw).Encode(h)
			return
		}
	}
	// json.NewEncoder(rw).Encode(&Horse{})

	// // return the horse
	// fmt.Fprintf(rw, "%v", horseName)
	// return
}

// Update horse name
func updateHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	horseName := params["name"]
	horseColor := params["color"]
	horseAge, _ := strconv.Atoi(params["age"])

	// check if the horse exists
	exist := ifHorseExist(horseName)

	if !exist {
		msg := fmt.Sprintf("Cannot update, Horse named %s does not exist", horseName)
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

	newHorse := Horse{Name: horseName, Color: horseColor, Age: horseAge}

	// newHorse := Horse{color = horseColor, name = horseName}

	// get the body of the request
	// body, err := ioutil.ReadAll(r.Body)
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		err := json.NewDecoder(r.Body).Decode(&newHorse)
		if err != nil {
			msg := fmt.Sprintf("Error decoding / reading json body - %v", err)
			fmt.Fprint(rw, msg)
			Check(err, msg)
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
		candidate  Horse
		better     Person
		investment int
		params     = mux.Vars(r)
	)

	// ***
	// get the params from the url
	ID := strconv.Itoa(rand.Intn(100))
	// id, _ := strconv.Atoi(params["id"])

	// par := fmt.Sprintf("params: %v \n", params)
	// fmt.Fprint(rw, par)

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
		fmt.Fprint(rw, msg)
		LogToFile(msg)
		return
	}

	// ***
	amountPositive := false

	// check if the amount is a number [or strconv.Atoi() the param (commented above)]
	// amountIsNumber := false
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
		// amountIsNumber = true
		amountPositive = true

		//
		if investment > (candidate.Record.Wins * 100) {
			msg := fmt.Sprintf("investment is -> %v, candidate record -> L: %v / W: %v", investment, candidate.Record.Losses, candidate.Record.Wins)
			expln := fmt.Sprintf("Amount must be proportional to %v's record (win / lose) odds", candidate.Name)

			fmt.Fprintf(rw, msg, expln)
			LogToFile(msg + expln)
			return
		}
	}

	if !amountPositive {
		msg := "Amount isnt a positive number, correct your request"
		// msg := "Amount isnt a number, correct your request"
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

			result, insertErr := conn.Collection.InsertOne(r.Context(), thisBet)
			Check(insertErr, "err on db objects insertion.")

			msg := fmt.Sprintf("Inserted a single document: %v", result)
			LogToFile(msg)
			// return the complet new bet
			json.NewEncoder(rw).Encode(thisBet)
			json.NewEncoder(rw).Encode(MainBoard)
		}
		// return the updated record
		fmt.Fprintf(rw, "%v", horses)
		fmt.Fprintf(rw, "%v", MainBoard)
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

// // removeBets remove the occurrences of an element from a slice, return the new slice
// func removeBets(slice []string, elem string) []string {
// 	// Create a new slice to store the result
// 	newSlice := make([]string, 0, len(slice))
// 	// Loop over the elements in the original slice
// 	for _, value := range slice {
// 		// If the element is not the one to remove, add it to the new slice
// 		if value != elem {
// 			newSlice = append(newSlice, value)
// 		}
// 	}
// 	return newSlice
// }

// func Status(rw http.ResponseWriter, r *http.Request) {
// 	rw.Write([]byte("API is up and running"))
// }
