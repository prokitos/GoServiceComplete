package app

import (
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Router /delete [get]
func deleteGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a delete request")

	param1 := r.FormValue("id")
	inter, err := strconv.Atoi(param1)
	if err != nil {
		log.Error("the delete request was not executed")
		log.Debug("id couldn't convert to a number: " + param1)
		w.Write([]byte(`"message": "Delete request failed"`))
		return
	}

	w.Write([]byte(`"message": "Delete request succes",`))
	services.DeleteDataEncrichment(w, inter)
}
