package app

import (
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Users godoc
// @Summary Delete persons by id from users
// @Description Delete persons by id from users
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query int true "delete user"
// @Failure 400 {string} string "{code:400, msg:"failure"}"
// @Success 200 {object} string "OK"
// @Router /delete [delete]
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
