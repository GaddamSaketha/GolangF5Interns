package mongodbF5interns

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

////////////////////////////////////////////////////////---------TRAINER STRUCT------------/////////////////////////////////////////////////////////////////////////////////////////////////
// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

///////////////////////////////////////////////////////-------DB_INTERFACE---------------/////////////////////////////////////////////////////////////////////////////////////////////
type connection interface {
	Connect()
}
type insertion interface {
	Insert(trainer Trainer)
}
type updation interface {
	Update()
}
type deletion interface {
	Delete()
}
type retreival interface {
	Retreive()
}

/////////////////////////////-------RETREIVE FUNCTION-------///////////////////////////////////////////////////
func (db *mongodb) Retreive() {
	var result Trainer
	filter := bson.D{{"name", "Ash"}}
	err := ((*db).collection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}

/////////////////////////////-----DELETE_FUNCTION//////////////////////////////////////////////////////////////
func (db *mongodb) Delete() {
	filter := bson.D{{"name", "Ash"}}
	deleteResult, err := ((*db).collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

/////////////////////////////////////////////////////////--------UPDATE_FUNCTION---------//////////////////////////////////////////////////////////////////
func (db *mongodb) Update() {
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := ((*db).collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

}

////////////////////////////////////////////////////////////---INSERT_FUNCTION------////////////////////////////////////////////////////////////////////////
func (db *mongodb) Insert(trainer Trainer) {
	collection := *db.collection
	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

//////////////////////////////////////////////////////////---CONNECT FUNCTION---------//////////////////////////////////////////////////////////////////////////////
func (db *mongodb) Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("trainers")
	fmt.Println("Connected to MongoDB!")
	(*db).client = client
	(*db).clientOptions = clientOptions
	(*db).collection = collection
}

///////////////////////////////////////////////////////////// --- MongoDB----/////////////////////////////////////////////////////////////////
type mongodb struct {
	client        *mongo.Client
	clientOptions *options.ClientOptions
	collection    *mongo.Collection
}

/////////////////////////////////////////////////////////////-----MAIN-----------///////////////////////////////////////////////////////////////////////
