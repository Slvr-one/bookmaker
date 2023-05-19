# This is a Go program that sets up a web server and defines several endpoints to handle HTTP requests. 

-- using the Gorilla mux package to handle routing and the net/http package to create the web 
-- server. The program defines several structs to represent data, including Horse, Record, Person, and Bet.

## The program defines several HTTP endpoints:

-- "/" - A welcome page that displays the number of horses available and the current date and time.
-- "/health" - A health check endpoint that returns a simple message to indicate that the server is up and running.
-- "/horses" - An endpoint that returns the number of horses available.
-- "/horses/{name}" - An endpoint that returns information about a specific horse by name.
-- "/horses/{name}/bet/{amount}" - An endpoint that allows the user to place a bet on a specific horse.
-- The program also defines several functions to handle these endpoints. The GetHorses function returns the number of horses available, the GetHorse function returns information about a specific horse, and the Invest function allows the user to place a bet on a specific horse. The main function initializes the web server and sets up the routing.

-----

##  initializing a MongoDB client, connect to a MongoDB database, and insert multiple documents into a collection/table:

-- The initMongoDB function initializes a MongoDB client with the URL "mongodb://localhost:27017", connects to the MongoDB database named "bookmaker", 
-- creates a collection named "bets" in the "bookmaker" database,
-- inserts three documents into the "bets" collection using the InsertMany method.
-- Each document is a BSON document that consists of several fields, such as:
    "fullName", 
    "age", 
    "amount", 
    "profit". 

-- The bson.D type is used to represent a BSON document in Go code. 
-- After inserting the documents, the InsertedIDs field of the InsertManyResult object is printed to the console.