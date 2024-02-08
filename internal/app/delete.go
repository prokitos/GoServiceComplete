package app

import (
	"modular/internal/models"
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

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Error("the delete request was not executed")
		log.Debug("id couldn't convert to a number: " + id)

		models.BadResponseSend(w, "operation failed, wrong id = "+id, 400)
		return
	}

	services.DeleteDataEncrichment(w, id)
}
