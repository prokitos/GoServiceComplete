package app

import (
	"modular/internal/models"
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Users godoc
// @Summary Delete persons by id from users
// @Description Delete persons by id from users
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "delete user"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /test/{id} [delete]
func delByIdRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a delete request")

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	vars := mux.Vars(r)
	id := vars["id"]
	if _, err := strconv.Atoi(id); err != nil {
		models.BadResponseSend(w, "operation failed, wrong id = "+id, 400)
		return
	}

	services.DeleteDataEncrichment(w, id)
}
