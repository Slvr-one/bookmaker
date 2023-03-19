package second

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// for monitor "OK" loop
func Ping() string {
	resp, _ := http.Get("localhost:5000/health")

	if resp != nil && resp.StatusCode == 200 {
		return "OK, Got Health!!"
	} else {
		return "Not Ok!"
	}
}

func initMongoDB() {
	var mongoURL string
	//init mongo client instance
	client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL)) //"mongodb://localhost:27017"
	if error != nil {
		//terminate app if errored on db connection
		panic(error)
	}

	//check if MongoDB database has been found and connected
	if error := client.Ping(context.TODO(), readpref.Primary()); error != nil {
		panic(error)
	}

	//init a collection (table) of bets in bookmaker db
	betsCollection := client.Database("bookmaker").Collection("bets")

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	bets := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}, {"amount", 250}, {"profit", 500}},
		bson.D{{"fullName", "User 3"}, {"age", 20}, {"amount", 550}, {"profit", 1100}},
		bson.D{{"fullName", "User 4"}, {"age", 28}, {"amount", 420}, {"profit", -110}},
	}
	// insert the bson object slice using InsertMany()
	results, error := betsCollection.InsertMany(context.TODO(), bets)
	// check for errors in the insertion
	if error != nil {
		panic(error)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
}

/////////////////////////////////////////////////////

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

func useCPU(w http.ResponseWriter, r *http.Request) {
	count := 1

	for i := 1; i <= 1000000; i++ {
		count = i
	}

	fmt.Printf("count: %d", count)
	w.Write([]byte(string(count)))
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

func printUniqueValue(arr []int) {
	//Create a dictionary for values of each element
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
}

////////////////////////////////////////////////////////////////////////////////

func second() {

	var tpl = template.Must(template.ParseFiles("static"))

	// Read the MongoDB username and password from the Secret
	username, err := getSecret("mongodb-credentials", "username")
	if err != nil {
		log.Fatal(err)
	}
	password, err := getSecret("mongodb-credentials", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB database using the username and password
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb:%v", mongoPath)))

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	initMongoDB()

	//Initialize an array
	inputArray := []int{10, 20, 30, 56, 67, 90, 10, 20}
	printUniqueValue(inputArray)

	router.HandleFunc("/search", searchHandler)
	router.HandleFunc("/user", userHandler)

	router.HandleFunc("/cpu", useCPU)

	//cach & return err
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
