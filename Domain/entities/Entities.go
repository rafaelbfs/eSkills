package entities

import (
	c "github.com/rafaelbfs/GoConvenience/Convenience"
	"github.com/rafaelbfs/MongoDBExtensions/dbexts"
	"github.com/rafaelbfs/eSkills/Domain/generated/pb/skills"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	t "time"
)

type Certification struct {
	Name       string
	Issuer     string
	Code       string
	Keywords   []string
	DateIssued t.Time
}

func CertificationFromPb(pb *skills.Certification) Certification {
	var origDate = pb.IssuedOn
	var publDate t.Time
	if origDate == nil {
		publDate = t.Now()
	} else {
		publDate = origDate.AsTime()
	}
	return Certification{Name: pb.Name, Issuer: pb.Issuer, Code: pb.Hash, DateIssued: publDate,
		Keywords: pb.Keywords}
}

func (c Certification) ToProto() *skills.Certification {
	var prot = skills.Certification{Name: c.Name, Issuer: c.Issuer}
	prot.Keywords = c.Keywords
	prot.Hash = c.Code
	return &prot
}

type Publication struct {
	Title           string
	PeriodicName    string
	Pages           string
	Url             string
	ISBN            string
	Keywords        []string
	PublicationDate t.Time
}

type ArticleParticipation struct {
	Participation skills.Publication_ParticipationType
	TheArticle    Publication
}

func (c ArticleParticipation) ToProto() *skills.Publication {
	var prot = skills.Publication{Title: c.TheArticle.Title, Periodic: c.TheArticle.PeriodicName,
		Participation: c.Participation}
	prot.Keywords = c.TheArticle.Keywords
	prot.Link = c.TheArticle.Url
	return &prot
}

func PublicationFromPb(pb *skills.Publication) ArticleParticipation {
	var p = ArticleParticipation{Participation: pb.Participation}
	p.TheArticle = Publication{Title: pb.Title, Keywords: pb.Keywords, PeriodicName: pb.Periodic,
		Pages: pb.Pages, Url: pb.Link}
	return p
}

type Experience struct {
	Position    string
	Company     string
	Description string
	StartDate   t.Time
	EndDate     *t.Time
	Keywords    []string
}

type OpenSourceProject struct {
	Repository  string
	LastCommit  t.Time
	Description string
	FirstCommit *t.Time
	Keywords    []string
}

type Person struct {
	ID             *primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      string
	LastName       string
	Certifications []Certification
	Articles       []ArticleParticipation
	WorkExperience []Experience
}

func PersonFromPb(pb skills.Person) Person {
	p := Person{FirstName: *pb.FirstName, LastName: *pb.LastName,
		Certifications: c.FMap(CertificationFromPb, pb.Experience.Certifications),
		Articles:       c.FMap(PublicationFromPb, pb.Experience.Publications)}

	if pb.Id != nil && dbexts.HexRegex.MatchString(*pb.Id) {
		p.ID = dbexts.ToOID(*pb.Id)
	}

	return p
}

func (p *Person) ToProto() *skills.Person {
	var prot = skills.Person{FirstName: &p.FirstName, LastName: &p.LastName}
	prot.Experience = new(skills.PersonAchievements)
	prot.Experience.Certifications = c.FMap(Certification.ToProto, p.Certifications)
	prot.Experience.Publications = c.FMap(ArticleParticipation.ToProto, p.Articles)
	if p.ID != nil && !p.ID.IsZero() {
		h := p.ID.Hex()
		prot.Id = &h
	}

	return &prot
}

func AppendId(person Person) bson.D {
	b := c.Try(bson.Marshal(person)).ResultOrPanic()
	var d bson.D
	err := bson.Unmarshal(b, &d)
	if err != nil {
		log.Panicf("Error! %v", err)
	}
	if person.ID != nil {
		d = append(d, bson.E{Key: "id", Value: person.ID.Hex()})
		return d
	}
	return d
}

func NewOid() *primitive.ObjectID {
	oid := primitive.NewObjectID()
	return &oid
}

type Skill struct {
	Name            string
	OwnRating       uint8
	OccurrenceRatio float64
}
