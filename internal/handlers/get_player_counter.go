package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diommyliel/GoAPI/api"
	"github.com/diommyliel/GoAPI/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetPlayerCounter(w http.ResponseWriter, r *http.Request) {
	var params = api.PlayerCounterParam{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return
	}

	var db *tools.DatabaseInterface
	db, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return
	}

	counterDetails := (*db).GetCounterDetails(params.Player)
	if counterDetails == nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return
	}

	var resp = api.PlayerCounterGetResponse{
		Code:    http.StatusOK,
		Counter: (*counterDetails).Counter,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return
	}

}
