package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	h "github.com/Slvr-one/bookmaker/handlers"
)

func SqlDB() *sql.DB {
	// logger, _ := thoth.Init("log")

	//checking env for db user & pass, and host (url)
	user, exist := os.LookupEnv("DB_USER")
	if !exist {
		h.LogToFile("DB_USER not set in .env")
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")
	if !exist {
		h.LogToFile("DB_PASS not set in .env")
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")
	if !exist {
		h.LogToFile("DB_HOST not set in .env")
		log.Fatal("DB_HOST not set in .env")
	}

	// setting credentials and trying a db connection
	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)
	db, sqlConnErr := sql.Open("mysql", credentials)
	// db, err := sql.Open("mysql",user+":"+pass+"@/"+host)
	h.Check(sqlConnErr, "Error on connection to sql db, credentials problem or likewise")
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

	h.Check(sqlConnErr, "Error on creating db todo, already exist?")

	_, sqlConnErr = db.Exec(`USE gotodo`)
	h.Check(sqlConnErr, "Error on setting gotodo as db")

	_, sqlConnErr = db.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)
	h.Check(sqlConnErr, "Error on creating table")

	return db
}
