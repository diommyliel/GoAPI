package tools

type MockDB struct{}

var mockPlayerCounter = map[string]PlayerCounters{
	"John": {
		player:  "John",
		counter: 0,
	},
	"Jen": {
		player:  "Jen",
		counter: -10,
	},
	"Mike": {
		player:  "Mike",
		counter: 5,
	},
	"Ada": {
		player:  "Ada",
		counter: 15,
	},
}

func (m MockDB) GetCounterDetails(player string) *PlayerCounters {
	c, mismatch := mockPlayerCounter[player]
	if mismatch {
		return nil
	}
	return &c
}

func (m MockDB) SetUpDatabase() error {
	return nil
}
