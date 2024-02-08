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
// @Summary Show persons in users
// @Description Show persons in users
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int false "Show users by id"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Success 200 {object} showUser "successful operation"
// @Router /test/{id} [get]
func showByIdRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a show request")

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	vars := mux.Vars(r)
	id := vars["id"]
	if _, err := strconv.Atoi(id); err != nil {
		models.BadResponseSend(w, "operation failed, wrong id = "+id, 400)
		return
	}

	var searchPerson models.User
	searchPerson.Id = id
	services.ShowSpecData(w, "", "", "", searchPerson)

}
