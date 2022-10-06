package templates

import (
	convenience "github.com/rafaelbfs/GoConvenience/Convenience"
	"github.com/rafaelbfs/eSkills/Domain/entities"
	ht "html/template"
	"log"
	"net/http"
)

const PERSON_DETAILS_PATH = "/page-fragments/person/details/"

var htmlTemplate *ht.Template
var dummy = entities.Person{FirstName: "John", LastName: "Doe"}

func errorHandler(err error) {
	if err != nil {
		log.Panicf("Could not load personDetails.html due to %v", err)
	}
}

func init() {
	var err error
	htmlTemplate, err = ht.ParseFiles("Frontend/templates/personDetails.html")
	convenience.WrapError(err).AndHandle(errorHandler)
}

func HandlePersonDetails(w http.ResponseWriter, r *http.Request) {
	extra := r.URL.Path[len(PERSON_DETAILS_PATH):]
	var per entities.Person
	if extra == "test" {
		per = dummy
	} else {
		per = entities.Person{FirstName: "New", LastName: "Person"}
	}

	htmlTemplate.Execute(w, per)
}

func GetPath(file string) string {
	return "Frontend/templates/" + file
}
