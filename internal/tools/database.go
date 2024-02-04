package tools

import log "github.com/sirupsen/logrus"

type PlayerCounters struct {
	Player  string
	Counter int
}

type DatabaseInterface interface {
	GetCounterDetails(player string) *PlayerCounters
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var db DatabaseInterface = &MockDB{}
	var err error = db.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
