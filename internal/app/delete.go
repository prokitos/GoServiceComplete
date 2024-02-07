package app

import (
	"encoding/json"
	postgres "modular/internal/models"
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
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /delete [delete]
func deleteGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a delete request")

	param1 := r.FormValue("id")
	inter, err := strconv.Atoi(param1)
	if err != nil {
		log.Error("the delete request was not executed")
		log.Debug("id couldn't convert to a number: " + param1)

		errResp := postgres.ErrorResponse{
			Message: "Invalid Input",
			Code:    400,
		}
		json.NewEncoder(w).Encode(errResp)

		return
	}

	services.DeleteDataEncrichment(w, inter)
}
