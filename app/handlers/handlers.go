package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	s "github.com/Slvr-one/bookmaker/app/structs"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/ichtrojan/thoth"
	"github.com/rs/zerolog/log"
	"rsc.io/quote"
)

var (
	Horses    []s.Horse
	MainBoard s.Board

	Conn    s.Conn
	connErr error

	// id        int
	// item      string
	// completed int
	// view     = template.Must(template.ParseFiles("./views/index.html"))
	// db = SqlDB()
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

	serverPort := ctx.GetString("serverPort")
	host := ctx.Request.Host
	port, _ := strconv.Atoi(serverPort) //r.URL.Query().Get("port")

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

///////!SECTION

func Check(err error, msg string) {
	if err != nil {
		color.Green(msg)
		LogToFile(msg)
		log.Info().Msg(msg)
		log.Fatal()

		// clog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		// clog.Fatalf("Error: %v\n Error Message: %v", err, msg)

		// panic(err.Error())
		// os.Exit(-1)
	}
}

func LogToFile(msg string) {
	// init a general file for logging
	genLogger, _ := thoth.Init("log")
	genLogger.Log(errors.New(msg))
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func RandomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Great to see you, %v!",
		"Look who it is! its %v",
		"Hi, %v. Welcome!",
		"Hi %v, knew you'd come back",
		"%v, You again..",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	i := rand.Intn(len(formats))
	return formats[i]
}

func EndProgram(start time.Time) {
	// log to file <end session> msg
	t := time.Now()
	elapsed := t.Sub(start)
	// elapsed = time.Since(start)
	msg := fmt.Sprintf("ended at -> %s, elapsed -> %s", t, elapsed)
	LogToFile(msg)
	// log.Info().Msg(msg)
}

func IfHorseExist(horseName string, horses []s.Horse) (exist bool) {
	for _, h := range horses {
		if h.Name == horseName {
			exist = true
			break
		}
		exist = false
	}
	return exist
}

func Populate() {
	// Hardcoded data - @todo: add database
	Horses = append(Horses, s.Horse{Name: "Monahen Boy", Color: "brown", Record: &s.Record{Wins: 8, Losses: 3}})
	Horses = append(Horses, s.Horse{Name: "Dangerous", Color: "brown:white", Record: &s.Record{Wins: 7, Losses: 1}})
	Horses = append(Horses, s.Horse{Name: "Black Beauty", Color: "black", Record: &s.Record{Wins: 4, Losses: 5}})
	Horses = append(Horses, s.Horse{Name: "horse 4", Color: "black", Record: &s.Record{Wins: 4, Losses: 5}})

}
