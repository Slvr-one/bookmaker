package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/propagation"
	// "go.opentelemetry.io/otel/trace"
)

// welcom - html
func Welcom(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	// ctx.Writer.Header().Set("Content-Type", "text/html")
	// ctx.Writer.Header().Set("Accept-Charset","utf-8")

	// ctx.Writer.Header().Set("Cache-Control", "no-cache")
	// ctx.Writer.Header().Set("Cache-Control","no-cache, no-store, must-revalidate")

	// ctx.Writer.Header().Set("Pragma","no-cache")
	// ctx.Writer.Header().Set("Expires","0")
	// ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
	// ctx.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// ctx.Writer.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
	// ctx.Writer.Header().Set("Access-Control-Expose-Headers","Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// ctx.Writer.Header().Set("Access-Control-Allow-Credentials","true")
	// ctx.Writer.Header().Set("Content-Encoding", "gzip")
	// ctx.Writer.Header().Set("Connection", "keep-alive")

	// fmt.Fprintf(ctx.Writer, Html)

	tpl := template.Must(template.ParseFiles("templates/index2.html"))
	dt := time.Now()
	// user :=

	goQoute := quote.Go()
	// message := fmt.Sprintf(randomFormat(), user)

	tpl.Execute(ctx.Writer, nil)
	// io.WriteString(ctx.Writer, "Hello again, Gefyra!")

	fmt.Fprintf(ctx.Writer, "<h3 style='color: black'>horses available: %v</h3>", len(horses))
	fmt.Fprintf(ctx.Writer, "<h3 style='color: black'> random go quote: %v</h3>", goQoute)
	fmt.Fprintf(ctx.Writer, "<h5 style='color: black'> %s</h5>", dt.Format("01-02-2006 15:04:05 Mon"))
}

// health - html
func Health(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/html")
	ctx.Writer.Write([]byte("<h3 style='color: steelblue'>Got Health</h3>"))

}

// func httpErrorBadRequest(err error, ctx *gin.Context) {
// 	httpError(err, ctx, http.StatusBadRequest)
// }

// func httpErrorInternalServerError(err error, ctx *gin.Context) {
// 	httpError(err, ctx, http.StatusInternalServerError)
// }

// func httpError(err error, ctx *gin.Context, status int) {
// 	log.Println(err.Error())
// 	ctx.String(status, err.Error())
// }

// func pingHandler(ctx *gin.Context) {
// 	req := resty.New().R().SetHeader("Content-Type", "application/text")
// 	otelCtx := ctx.Request.Context()
// 	span := trace.SpanFromContext(otelCtx)
// 	defer span.End()
// 	otel.GetTextMapPropagator().Inject(otelCtx, propagation.HeaderCarrier(req.Header))
// 	url := ctx.Query("url")
// 	if len(url) == 0 {
// 		url = os.Getenv("PING_URL")
// 		if len(url) == 0 {
// 			httpErrorBadRequest(errors.New("url is empty"), ctx)
// 			return
// 		}
// 	}
// 	log.Printf("Sending a ping to %s", url)
// 	resp, err := req.Get(url)
// 	if err != nil {
// 		httpErrorBadRequest(err, ctx)
// 		return
// 	}
// 	log.Println(resp.String())
// 	ctx.String(http.StatusOK, resp.String())
// }

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

func Monitor(ctx *gin.Context) {

	// port :=
	host := ctx.Request.Host
	port, _ := strconv.Atoi(defaultServerPort) //r.URL.Query().Get("port")

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
	json.NewEncoder(ctx.Writer).Encode(horses)
	// fmt.Fprintf(ctx.Writer, "%v", horses)

	// hLen := len(horses)
	// json.NewEncoder(ctx.Writer).Encode(hLen)
}

func CreateHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	exist := ifHorseExist(horseName)

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
			// append() adds elements to a slice, but in this case, its used to remove an element:
			// The [:i] notation denotes that all the elements before the i-th index are included,
			// while [i+1:] denotes that all the elements after the i-th index are included.
			// By combining these two notations with the ellipsis(...), the horses slice is effectively modified to exclude the i-th element.
			// note that the original slice is not modified; rather, a new slice is created missing the i-th element.
			horses = append(horses[:i], horses[i+1:]...)
			break
		}
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

// fetch horse by name
// ctx.Writer http.ResponseWriter, r *http.Request
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

// func userHandler(ctx.Writer http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	if id == "" {
// 		http.Error(ctx.Writer, "The id query parameter is missing", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Fprintf(ctx.Writer, "<h1>The user id is: %s</h1>", id)
// }

// func searchHandler(ctx.Writer http.ResponseWriter, r *http.Request) {
// 	u, err := url.Parse(r.URL.String())

// 	if err != nil {
// 		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
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
// // func updateHorse(ctx *gin.Context) {

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
// 	return http.HandlerFunc(func(ctx *gin.Context) {
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

// func Status(ctx *gin.Context) {
// 	ctx.Writer.Write([]byte("API is up and running"))
// }
