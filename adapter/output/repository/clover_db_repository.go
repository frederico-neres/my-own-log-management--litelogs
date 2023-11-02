package repository

import (
	"fmt"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/domain"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/port"
	c "github.com/ostafen/clover"
)

type cloverDbRepository struct {
	collection string
}

func NewCloverDbRepository() port.LogRepository {
	collection := "logCollection"
	db, err := c.Open("clover-db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = db.CreateCollection(collection)
	if err != nil {
		fmt.Println(err)
	}

	return &cloverDbRepository{
		collection: collection,
	}
}

func (r *cloverDbRepository) Save(logDomain *domain.LogDomain) error {
	db, err := c.Open("clover-db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	doc := c.NewDocument()
	doc.Set("time", logDomain.Time)
	doc.Set("level", logDomain.Level)
	doc.Set("service", logDomain.Service)
	doc.Set("message", logDomain.Message)

	docId, _ := db.InsertOne(r.collection, doc)
	fmt.Println(docId)

	return nil
}
