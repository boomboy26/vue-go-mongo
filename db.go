package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func db() *mongo.Client {

clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")// Connect to //MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
log.Fatal(err)
}
// Check the connection
err = client.Ping(context.TODO(), nil)
if err != nil {
log.Fatal(err)
}
fmt.Println("Connected to MongoDB!")
return client
}

type user struct{
	Name string `json:"name"`
	City string `json:"city"`
	Age int `json:"age"`
	}
	
func createProfile(w http.ResponseWriter, r *http.Request) {w.Header().Set("Content-Type", "application/json") // for adding       //Content-type
var person user
var userCollection = db().Database("goTest").Collection("users")
err := json.NewDecoder(r.Body).Decode(&person) // storing in person   //variable of type user
if err != nil {
fmt.Print(err)
}
insertResult, err := userCollection.InsertOne(context.TODO(),person)
if err != nil {
log.Fatal(err)
}
json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
}