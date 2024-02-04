package tools

type MockDB struct{}

var mockPlayerCounter = map[string]PlayerCounters{
	"John": {
		Player:  "John",
		Counter: 0,
	},
	"Jen": {
		Player:  "Jen",
		Counter: -10,
	},
	"Mike": {
		Player:  "Mike",
		Counter: 5,
	},
	"Ada": {
		Player:  "Ada",
		Counter: 15,
	},
}

func (m *MockDB) GetCounterDetails(player string) *PlayerCounters {
	var c PlayerCounters
	c, ok := mockPlayerCounter[player]
	if !ok {
		return nil
	}
	return &c
}

func (m *MockDB) SetUpDatabase() error {
	return nil
}
