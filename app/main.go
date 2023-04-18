package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

// http://localhost:9000/invest/%7BDangerous%7D/%7B500%7D

type Record struct {
	Wins  int `json:"wins"`
	Loses int `json:"loses"`
}

type Horse struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Record *Record `json:"record"`
}

type Person struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname"  bson:"lasttname"`
}

type Bet struct {
	ID       string  `json:"id"       bson:"_id"`
	Amount   uint    `json:"amount"   bson:"amount"`
	Profit   int     `json:"profit"   bson:"profit"`
	Investor *Person `json:"investor" bson:"investor"`
}

// type DatabaseCollections struct {
// 	Bets *mongo.Collection
// }

// Init horses var as a slice of Horse struct (slice: an array with no fixed Size or Type)
var (
	horses []Horse
	bets   []Bet

	// serverPort = 5000
	// mongoPort  = 27017
)

// func init() {
// 	// Connect to the MongoDB database using the username and password

// 	//get port env var (if provided by user, default / 5000)
// 	mongoPort := os.Getenv("mongoPort")
// 	if mongoPort == "" {
// 		mongoPort = "27017"
// 	}

// 	mongodbUrl := fmt.Sprintf("mongodb://localhost:%s", mongoPort)

// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()

// 	//init mongo client
// 	clientOptions := options.Client().ApplyURI(mongodbUrl)
// 	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb:%v", mongoPath))

// 	client, connectionErr := mongo.Connect(ctx, clientOptions)
// 	// client, connectionErr := mongo.NewClient(clientOptions)

// 	// terminate app if errored on db connection
// 	HandleErr(connectionErr, "err on db connection.")

// 	//check if MongoDB database has been found and connected
// 	pingErr := client.Ping(ctx, readpref.Primary())
// 	HandleErr(pingErr, "err on db ping test.")

// 	//init a collection (table) of bets in bookmaker db
// 	bmDB := client.Database("bookmaker")
// 	betsCollection := bmDB.Collection("bets")

// 	// insert multiple documents into a collection
// 	// by createing a slice of bson.D objects
// 	bets := []interface{}{
// 		bson.D{{"fullName", "User 1"}, {"age", 30}, {"amount", 500}, {"profit", -100}},
// 		bson.D{{"fullName", "User 2"}, {"age", 25}, {"amount", 250}, {"profit", 500}},
// 		bson.D{{"fullName", "User 3"}, {"age", 20}, {"amount", 550}, {"profit", 1100}},
// 		bson.D{{"fullName", "User 4"}, {"age", 28}, {"amount", 420}, {"profit", -110}},
// 	}
// 	// insertMany returns a result object with the ids of the newly inserted objects
// 	result, insertErr := betsCollection.InsertMany(ctx, bets)
// 	HandleErr(insertErr, "err on db objects insertion.")

// 	// display the ids of the newly inserted objects
// 	// fmt.Println("Inserted a single document: ", result.InsertedID)
// 	fmt.Println("Inserted a single document: ", result)
// }

// main
func main() {
	// init a router handler for the server client
	router := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static"))

	//get port env var (if provided by user, default / 9090)
	serverPort := os.Getenv("serverPort")
	if serverPort == "" {
		serverPort = "9090"
	}

	listenAddr := flag.String("listenaddr", serverPort, "port to serve the app")

	// // checking for a local .env file containing vars - redundant as of now
	// envLoadErr := godotenv.Load()
	// helpers.HandleErr(envLoadErr, "err loading .env file.")

	// init a server client with custom spec - for listen & serve
	mainServer := &http.Server{
		Addr:           ":" + *listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // "1 times 2, 20 times" or 1048576 - standard size of header :)
		// BaseContext: func(l net.Listener) context.Context {
		// 	ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
		// 	return ctx
		// },

		// TLSConfig: &tls.Config{
		// 	ClientAuth: tls.RequireAndVerifyClientCert,
		// 	ClientCAs:  certPool,
		// 	MinVersion: tls.VersionTLS12,
		// 	CipherSuites: []uint16{
		// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		// 	}
		// 	PreferServerCipherSuites: true,
		// 	CurvePreferences: []tls.CurveID{
		// 		tls.CurveP256,
		// 		tls.X25519,
		// 		tls.CurveP384,

		// 	}
		// // InsecureSkipVerify: true,
	}

	// Hardcoded data - @todo: add database
	horses = append(horses, Horse{Name: "Monahen boy", Color: "brown", Record: &Record{Wins: 8, Loses: 3}})
	horses = append(horses, Horse{Name: "Dangerous", Color: "brown:white", Record: &Record{Wins: 7, Loses: 1}})
	horses = append(horses, Horse{Name: "horse 3", Color: "black", Record: &Record{Wins: 4, Loses: 5}})

	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")

	router.HandleFunc("/", Welcom).Methods("GET")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/metrics", Monitor).Methods("GET")

	router.HandleFunc("/LH", GetHorses).Methods("GET")             //List all available horses
	router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")       //Get a specific horse
	router.HandleFunc("/UH/{name}", updateHorse).Methods("UPDATE") //Update a specific horse
	router.HandleFunc("/invest/{horse}/{amount}", Invest).Methods("GET")

	log.Printf("set to listen on port: %v", serverPort)

	// log.Fatal(http.ListenAndServe(":"+serverPort, router))
	// log.Fatal(http.ListenAndServe(":"+*listenAddr, router))
	log.Fatal(mainServer.ListenAndServe())

}

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %s", msg)
		panic(err.Error())

	}
}

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
	// rw.Header().Set("Content-Type", "application/json")

	// hLen := len(horses)
	// json.NewEncoder(rw).Encode(hLen)

	// return the list of horses
	fmt.Fprintf(rw, "%v", horses)
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

	if !horseExists {
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
	json.NewEncoder(rw).Encode(&Horse{})

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

	// newHorse := types.Horse{color = horseColor, name = horseName}

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

	var (
		candidate  Horse
		investment int
		better     Person
	)
	// get the params from the url
	params := mux.Vars(r)
	candidate.Name = params["horse"]
	betterName := params["investor"]
	better.FirstName = betterName[:]
	better.FirstName = betterName[:]

	// investment, convertionErr := strconv.Atoi(params["amount"])
	investment, _ = strconv.Atoi(params["amount"])
	ID := strconv.Itoa(rand.Intn(100))
	// HandleErr(convertionErr, "Error during str conversion to int.")

	// check if the horse exists
	horseExists := false
	for _, h := range horses {
		if h.Name == candidate.Name {
			horseExists = true
			break
		}
	}

	if !horseExists {
		fmt.Fprintf(rw, "Horse %s does not exist", candidate.Name)
		return
	}

	// check if the amount is a number or strconv.Atoi() the param (commented above)
	amountIsNumber := false
	amountInt, err := fmt.Sscanf(string(investment), "%d")
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

	if !amountIsNumber {
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
		if h.Name == candidate.Name {
			quotient := h.Record.Loses / h.Record.Wins // https://www.cuemath.com/numbers/quotient/
			// odds := math.Round(quotient)

			// bet := types.Bet{Horse: horseName, Amount: amountInt}
			var bet Bet
			bet.ID = ID
			bet.Amount = uint(investment)
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
