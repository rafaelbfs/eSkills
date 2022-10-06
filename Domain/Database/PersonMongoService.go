package Database

import (
	c "github.com/rafaelbfs/GoConvenience/Convenience"
	"github.com/rafaelbfs/MongoDBExtensions/dbexts"
	"github.com/rafaelbfs/eSkills/Domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var personRepo *PersonRepository

func GetPersonRepo() *PersonRepository {
	if personRepo == nil {
		personRepo = MkPersonRepo()
	}
	return personRepo
}

type PersonRepository struct {
	repo *dbexts.MongoCRUDRepo[entities.Person]
}

func MkPersonRepo() *PersonRepository {
	dbexts.Initialize("local.env", "skillsdatabase")
	r := dbexts.MongoCRUDRepo[entities.Person]{}
	p := PersonRepository{repo: &r}
	p.Init()
	return &p
}

func (it *PersonRepository) Init() *PersonRepository {
	it.repo.Init("people")
	return it
}

func GetPersonRepository() *PersonRepository {
	if c.Nvl(personRepo).IsNil() {
		personRepo = MkPersonRepo()
	}
	return personRepo
}

func mkErrorHandler(op string) func(err error) {
	return func(err error) {
		log.Printf("Could not %v person due to %v", op, err)
	}
}

func (r *PersonRepository) InsertPerson(person *entities.Person) *mongo.InsertOneResult {
	res, err := r.Init().repo.Create(*person)
	c.WrapError(err).AndHandle(mkErrorHandler("create"))
	return res
}
