// package main

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"gopkg.in/mgo.v2/bson"
// )

// func TestCreateBet(t *testing.T) {
// 	// Arrange
// 	bet := Bet{
// 		Investor.:    "My Bet",
// 		Amount:  100,
// 		HorseID: 1,
// 	}

// 	// Act
// 	err := db.Save(&bet)

// 	// Assert
// 	if err != nil {
// 		t.Errorf("Error saving bet: %v", err)
// 	}

// 	// mongo

// 	db := client.Database("test_db")

//  	// Create a new user
//  	user := User{
//  	 FirstName: "John",
//  	 LastName:  "Doe",
//  	 Email:     "john.doe@example.com",
//  	}
//  	err = user.Create(context.Background(), db, "users", &user)
//  	if err != nil {
//  	 panic(err)
//  	}
//  	fmt.Printf("User created: %v\n", user)

//  	// Read a user by ID
//  	var readUser User
//  	err = readUser.Read(context.Background(), db, "users", bson.M{"_id": user.ID}, &readUser)
//  	if err != nil {
//  	 panic(err)
//  	}
//  	fmt.Printf("User read: %v\n", readUser)

//  	// Update a user's email
//  	update := bson.M{"$set": bson.M{"email": "john.doe_updated@example.com", "updated_at": primitive.NewDateTimeFromTime(user.UpdatedAt)}}
//  	err = user.Update(context.Background(), db, "users", bson.M{"_id": user.ID}, update)
//  	if err != nil {
//  	 panic(err)
//  	}
//  	fmt.Printf("User updated: %v\n", user)

//  	// Delete a user by ID
//  	err = user.Delete(context.Background(), db, "users", bson.M{"_id": user.ID})
//  	if err != nil {
//  	 panic(err)
//  	}
//  	fmt.Println("User deleted")
// }
