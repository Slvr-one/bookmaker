package main

import (
	"encoding/json"
	"fmt"
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

	"src/main.go/types"

	"github.com/gorilla/mux"
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

func Monitor(rw http.ResponseWriter, r *http.Request) {
	//if app is up, keep logging every 5
	for {
		// result := Ping()
		// fmt.Println(result)
		fmt.Printf("\r%s -- %d", "logs", rand.Intn(100000000)) //strconv.Itoa(rand.Intn(100000000)))
		time.Sleep(5 * time.Second)
	}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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
	json.NewEncoder(rw).Encode(&types.Horse{})
	// json.NewEncoder(rw).Encode(horses)
}

// Update horse name
func updateHorse(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range horses {
		if item.Name == params["name"] {
			horses = append(horses[:index], horses[index+1:]...) //redundent
			var horse types.Horse
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
			quotient := h.Record.Loses / h.Record.Wins // https://www.cuemath.com/numbers/quotient/
			// odds := math.Round(quotient)
			var bet types.Bet
			bet.ID = strconv.Itoa(rand.Intn(100))
			bet.Amount = uint(investment)
			bet.Profit = investment * quotient
			bet.Investor.FirstName = "dvir"
			bet.Investor.LastName = "gross"

			Bets = append(Bets, bet) //send to db

			json.NewEncoder(rw).Encode(bet)
			return
		}
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
