package iform

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ItsElijahWood/Goproject/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func IForm(w http.ResponseWriter, r *http.Request) {
	if database.Client == nil { 
		http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	age := r.FormValue("age")

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	collection := database.Client.Database("Forms").Collection("iForm") 
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	doc := bson.M{"name": name, "email": email, "age": ageInt}
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		http.Error(w, "Failed to insert data into MongoDB", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Data successfully inserted into MongoDB")
}
