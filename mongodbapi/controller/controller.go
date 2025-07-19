package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const conn = "mongodb://localhost:27017/"
const dbname = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {

	clientoption := options.Client().ApplyURI(conn)

	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection successfull")

	collection = client.Database(dbname).Collection(colName)

	fmt.Println("collecion instance is ready")
}

func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted id : ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	// bson.M = map[string]interface{}
	// Unordered key-value pairs.
	// Used in most places for filters and updates because it’s convenient.

	filter := bson.M{"_id": id}  

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deletecount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count : ", deletecount.DeletedCount)

}

func deleteAllMovies()  *mongo.DeleteResult {

	// bson.D = []bson.E
	// Ordered key-value pairs.
	// Required in some cases (like aggregation pipelines or when order matters).

	filter := bson.D{{}} //using empty "{}" means everything will be included. i.e all will be deleted

	deletedcount, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted all data : ", deletedcount.DeletedCount)

	return deletedcount
}

func getAllMovies() []primitive.M{

	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	
	var movies []primitive.M
	// var movies []primitive.M
	// What it is:
	// A slice of BSON documents (each is primitive.M, i.e., map[string]interface{}).
	// Used to store multiple documents returned by Find.
	// Why primitive.M instead of a struct:
	// Flexible, no need to define all fields beforehand.
	// If you’re unsure of structure or have dynamic schemas, it's easier to use primitive.M.


	// cur is a cursor:
	// Returned by collection.Find(...), it lazily loads documents from MongoDB.
	// Think of it as a pointer that goes over each result one by one.

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}


func GetAllMovies(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")

	allmovies:=getAllMovies()

	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")

	w.Header().Set("Allow-Control-Allow-Methods","POST")  //to allow only post methods

	var movie model.Netflix

	_=json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	
	json.NewEncoder(w).Encode(map[string]string{"message": "successfull"})
}


func MarkAsWatched(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")

	w.Header().Set("Allow-Control-Allow-Methods","PUT")  //to allow only put methods

	params:=mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params:=mux.Vars(r)

	deleteOneMovie(params["id"])
	fmt.Println(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count:=deleteAllMovies()


	json.NewEncoder(w).Encode(count)
}	