package rest

import (
	"fmt"
	"github.com/rafaelbfs/eSkills/Domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func HandlePersonRequests(w http.ResponseWriter, r *http.Request) {
	people := shouldMock(r)
	w.Header().Set("Content-Type", "application/json")

	j, err := bson.MarshalExtJSON(bson.D{{"people", people}}, false, true)
	if err != nil {
		fmt.Errorf("Error!\n %v", err)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
		w.WriteHeader(500)
		return
	}
	w.Write(j)
}

func shouldMock(r *http.Request) []entities.Person {
	if r.Method == "GET" && strings.Contains(r.Header.Get("x-mock"), "yes") {
		return []entities.Person{
			{FirstName: "John", LastName: "Doe", ID: entities.NewOid()},
			{FirstName: "Jane", LastName: "Doe", ID: entities.NewOid()},
			{FirstName: "John", LastName: "Davis", ID: entities.NewOid()},
			{FirstName: "Amanda", LastName: "Davis", ID: entities.NewOid()},
		}
	}
	return []entities.Person{}
}
