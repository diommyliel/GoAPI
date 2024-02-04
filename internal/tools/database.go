package tools

import "log"

type PlayerCounters struct {
	player  string
	counter int
}

type DatabaseInterface interface {
	GetCounterDetails(player string) *PlayerCounters
	SetUpDatabase() error
}

func NewDatabase(db DatabaseInterface) (*DatabaseInterface, error) {

	var err error = db.SetUpDatabase()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &db, nil
}
