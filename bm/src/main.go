package bookmaker

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// diclare json structs - horses
type Record struct {
	Wins  int `json:"wins"`
	Loses int `json:"loses"`
}

type Horse struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Record *Record `json:"record"`
}

// Init horses var as a slice of Horse struct (slice: an array with no fixed Size or Type)
var horses []Horse

///////////////////////////////////////////////////

type Person struct {
	firstName string `json:"firstname" bson:"firstname"`
	lastName  string `json:"lastname"  bson:"lasttname"`
}

type Bet struct {
	ID       string  `json:"id"       bson:"_id"`
	Amount   uint    `json:"amount"   bson:"amount"`
	Profit   int     `json:"profit"   bson:"profit"`
	Investor *Person `json:"investor" bson:"investor"`
}

var Bets []Bet

// var tpl = template.Must(template.ParseFiles("static/index.html"))

// welcom - html
func Welcom(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	// tpl.Execute(rw, nil)

	dt := time.Now()

	fmt.Fprintf(rw, "<h3 style='color: maroon'>horses available: %v</h3>", len(horses))
	fmt.Fprintf(rw, "<h5 style='color: black'> %s</h5>", dt.Format("01-02-2006 15:04:05 Mon"))
}

// health - html
func Health(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h3 style='color: steelblue'>Got Health</h3>"))
}

func Monitor(rw http.ResponseWriter, r *http.Request) {
	//if app is up, keep logging every 5
	for {
		// result := Ping()
		// fmt.Println(result)
		fmt.Printf("\r%s -- %d", "logs", rand.Intn(100000000)) //strconv.Itoa(rand.Intn(100000000)))
		time.Sleep(5 * time.Second)
	}
}

// fetch horses available (number)
func GetHorses(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var hLen = len(horses)
	json.NewEncoder(rw).Encode(hLen)
}

// fetch horse by name
func GetHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through horses, find one with name to mach "name"
	for _, item := range horses {
		if item.Name == params["name"] {
			// var horse Horse
			json.NewEncoder(rw).Encode(item)
			return
		}
	}
	json.NewEncoder(rw).Encode(&Horse{})
	// json.NewEncoder(rw).Encode(horses)
}

// Update horse name
func updateHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range horses {
		if item.Name == params["name"] {
			horses = append(horses[:index], horses[index+1:]...) //redundent
			var horse Horse
			_ = json.NewDecoder(r.Body).Decode(&horse)
			horse.Name = params["name"]
			horses = append(horses, horse)
			json.NewEncoder(rw).Encode(horse)
			return
		}
	}
}

// bet on a horse
func Invest(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	candidate := params["horse"]
	investment, err := strconv.Atoi(params["amount"])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	for _, h := range horses {
		if h.Name == candidate {
			quotient := h.Record.Loses / h.Record.Wins
			// odds := math.Round(quotient)
			var bet Bet
			bet.ID = strconv.Itoa(rand.Intn(100))
			bet.Amount = uint(investment)
			bet.Profit = investment * quotient
			bet.Investor.firstName = "dvir"
			bet.Investor.lastName = "gross"

			Bets = append(Bets, bet) //send to db

			json.NewEncoder(rw).Encode(bet)
			return
		}
	}
}

// main
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//get env for port (if provided by user, 5000 is default)
	port := os.Getenv("port")
	if port == "" {
		port = "5000"
	}

	router := mux.NewRouter()
	// mux := http.NewServeMux()

	// Hardcoded data - @todo: add database
	horses = append(horses, Horse{Name: "Monahen boy", Color: "brown", Record: &Record{Wins: 8, Loses: 3}})
	horses = append(horses, Horse{Name: "Dangerous", Color: "brown:white", Record: &Record{Wins: 7, Loses: 1}})
	horses = append(horses, Horse{Name: "horse 3", Color: "black", Record: &Record{Wins: 4, Loses: 5}})

	// fs := http.FileServer(http.Dir("static"))

	// router.Handle("/", fs) //.Methods("GET")
	// router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// router.Handle("/assets/", http.StripPrefix("/assets/", staticHandler()).Methods("GET")

	router.HandleFunc("/home", Welcom).Methods("GET")
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/metrics", Monitor).Methods("GET")

	router.HandleFunc("/LH", GetHorses).Methods("GET")          //List all available horses
	router.HandleFunc("/GH/{name}", GetHorse).Methods("GET")    //Get a specific horse
	router.HandleFunc("/UH/{name}", updateHorse).Methods("PUT") //Update a specific horse
	// router.HandleFunc("/invest/{horse}/{amount}", Invest).Methods("UPDATE")

	log.Printf("Listening on port: %v...", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
