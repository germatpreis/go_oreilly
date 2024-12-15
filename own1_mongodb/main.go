package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	// https://www.mongodb.com/docs/drivers/go/current/quick-start/#std-label-golang-quickstart
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("Set your 'MONGDOB_URI' env variable")
	}

	log.Printf("Trying to connect to mongo using uri %s", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	log.Printf("Connection established")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	databaseName := "sample_mflix"
	collectionName := "movies"

	database := client.Database(databaseName)

	ensureCollectionExists(database, collectionName)

	coll := database.Collection("movies")

	title := "Back to the Future"

	log.Printf("Trying to read from collection %s ", coll.Name())

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)

	log.Printf("Read from collection %s, lets see what happens", coll.Name())

	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the title %s\n", title)
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}

func ensureCollectionExists(database *mongo.Database, collectionName string) {
	_, err := database.ListCollectionNames(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	//if err != nil {
	//	log.Printf("Collection %s does not exist, creating it")
	//	err := database.CreateCollection(context.TODO(), collectionName)
	//	if err != nil {
	//		log.Fatal("Couldn't create collection", collectionName)
	//	}
	//	log.Printf("Created collection %s", collectionName)
	//}
}
