package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// mongo

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//sql
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	// "github.com/ichtrojan/thoth"
)

var (
// id        int
// item      string
// completed int
// view     = template.Must(template.ParseFiles("./views/index.html"))
// db = SqlDB()
)

func MongoDB(mongodbUrl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//init mongo client
	clientOptions := options.Client().ApplyURI(mongodbUrl)
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb:%v", mongoPath))

	// conn = &Conn{}
	// conn.Client, err = mongo.Connect(ctx, clientOptions)
	client, connErr := mongo.Connect(ctx, clientOptions)
	// client, connErr := mongo.NewClient(clientOptions)
	// Check(connErr, "err on db client creation.")
	// connErr = client.Connect(ctx)
	Check(connErr, "err on db connection.")

	defer func() {
		connErr := client.Disconnect(ctx)
		Check(connErr, "db dissconnected")
	}()

	//check if MongoDB database has been found and connected
	pingErr := client.Ping(ctx, readpref.Primary())
	Check(pingErr, "err on db ping test")

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// Check(err, "err on listing databases")
	// fmt.Println(databases)

	//init a collection (table) of bets in bookmaker db
	bmDB := client.Database("bookmaker")
	betsColl := bmDB.Collection("bets")
	conn.Collection = betsColl

	MainBoard.Bets[0].Investor.UserName,
		MainBoard.Bets[1].Investor.UserName,
		MainBoard.Bets[2].Investor.UserName,
		MainBoard.Bets[3].Investor.UserName =
		"User 1", "dviross", "root", "Fat Guy"

	MainBoard.Bets[0].Investor.Age, MainBoard.Bets[1].Investor.Age, MainBoard.Bets[2].Investor.Age, MainBoard.Bets[3].Investor.Age = 30, 99, 52, 71
	MainBoard.Bets[0].Amount, MainBoard.Bets[1].Amount, MainBoard.Bets[2].Amount, MainBoard.Bets[3].Amount = 500, 70, 9900, 15
	MainBoard.Bets[0].Profit, MainBoard.Bets[1].Profit, MainBoard.Bets[2].Profit, MainBoard.Bets[3].Profit = -100, 400, -350, 1110
	// insert multiple documents into a collection
	// by createing a slice of bson.D objects
	var testBoard = []interface{}{
		bson.M{"fullName": "User 1", "age": 30, "amount": 1110, "profit": -100},
		bson.M{"fullName": "dviross", "age": 76, "amount": 78, "profit": 6900},
		bson.M{"fullName": "root", "age": 13, "amount": 2480, "profit": 1100},
		bson.M{"fullName": "Fat Guy", "age": 84, "amount": 30500, "profit": -13000},
	}

	// insertMany returns a result object with the ids of the newly inserted objects
	result, insertErr := betsColl.InsertMany(ctx, testBoard)
	Check(insertErr, "err on db objects insertion.")

	// display the ids of the newly inserted objects
	// fmt.Println("Inserted a single document: ", result.InsertedID)
	fmt.Println("Inserted a single document: ", result)

	// title := "Back to the Future"

	// var result bson.M
	// err = betsColl.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	msg := fmt.Sprintf("No document was found with the title %s\n", title)
	// 	Check(err, msg)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	return client
}

func SqlDB() *sql.DB {
	// logger, _ := thoth.Init("log")

	//checking env for db user & pass, and host (url)
	user, exist := os.LookupEnv("DB_USER")
	if !exist {
		LogToFile("DB_USER not set in .env")
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")
	if !exist {
		LogToFile("DB_PASS not set in .env")
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")
	if !exist {
		LogToFile("DB_HOST not set in .env")
		log.Fatal("DB_HOST not set in .env")
	}

	// setting credentials and trying a db connection
	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)
	db, sqlConnErr := sql.Open("mysql", credentials)
	// db, err := sql.Open("mysql",user+":"+pass+"@/"+host)
	Check(sqlConnErr, "Error on connection to sql db, credentials problem or likewise")
	defer db.Close()

	fmt.Println("Database Connection Successful - ⛁ -")

	// results, err := db.Query("SELECT user_id,username,password,first_name,last_name,dob,address FROM users")
	// Check(err, "Error on db query")

	// for results.Next() {
	// 	err := results.Scan(&user.Id,&user.Username,&user.Password,&user.FirstName,&user.LastName,&user.DOB,&user.Address)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	allUsers = append(allUsers, user)
	// }
	// _, sqlConnErr = db.Exec(`CREATE DATABASE gotodo`)

	Check(sqlConnErr, "Error on creating db todo, already exist?")

	_, sqlConnErr = db.Exec(`USE gotodo`)
	Check(sqlConnErr, "Error on setting gotodo as db")

	_, sqlConnErr = db.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)
	Check(sqlConnErr, "Error on creating table")

	return db
}
