package entities

import (
	assertions "github.com/rafaelbfs/GoConvenience/Assertions"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestAppendId(t *testing.T) {
	i := primitive.NewObjectID()
	john := Person{FirstName: "John", LastName: "Doe", ID: &i}
	d := AppendId(john)
	assertions.Assert(t).Condition(len(d) > 0).IsTrueV()
}
