package api

// Endpoint parms
type PlayerCounterParam struct {
	player string
}

// Get counter response
type PlayerCounterGetResponse struct {
	code    int
	counter int
}

// Error response
type Error struct {
	code    int
	message string
}
