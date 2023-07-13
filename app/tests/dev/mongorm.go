// package main

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func (b *Bet) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
// 	collection := db.Collection(collectionName)

// 	b.CreatedAt = time.Now()
// 	b.UpdatedAt = time.Now()

// 	res, err := collection.InsertOne(ctx, model)
// 	Check(err, "insertion err on create")

// 	b.ID = res.InsertedID.(primitive.ObjectID)
// 	return nil
// }

// func (b *Bet) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
// 	collection := db.Collection(collectionName)

// 	err := collection.FindOne(ctx, filter).Decode(result)
// 	Check(err, "insertion err on read")

// 	return nil
// }

// func (b *Bet) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
// 	collection := db.Collection(collectionName)

// 	b.UpdatedAt = time.Now()

// 	_, err := collection.UpdateOne(ctx, filter, update)
// 	Check(err, "insertion err on update")

// 	return nil
// }

// func (b *Bet) Delete(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
// 	collection := db.Collection(collectionName)
// 	_, err := collection.DeleteOne(ctx, filter)
// 	Check(err, "insertion err on delete")

// 	return nil
// }
