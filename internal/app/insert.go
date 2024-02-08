package app

import (
	"encoding/json"
	"io"
	"modular/internal/models"
	postgres "modular/internal/models"
	services "modular/internal/services/myService"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Users godoc
// @Summary Insert persons into users
// @Description Insert persons into users
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body addUser true "Add user"
// @Success 200 "successful operation"
// @Router /insert [post]
func insertGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a create request")

	reqBody, _ := io.ReadAll(r.Body)
	var user postgres.User
	json.Unmarshal(reqBody, &user)

	if len(user.Name) > 40 || len(user.Surname) > 40 || len(user.Patronymic) > 40 || len(user.Name) == 0 || len(user.Surname) == 0 || len(user.Patronymic) == 0 {
		log.Error("the insert request was not executed")
		log.Debug("incorrect length of persons data: " + "name=" + user.Name + " surname=" + user.Surname + "patronymic=" + user.Patronymic)

		models.BadResponseSend(w, "operation failed, wrong users param", 400)
		return
	}

	services.CreateDataEncrichment(w, user)
}

// для вывода в сваггер
type addUser struct {
	Name       string `json:"name" example:"ivan"`
	Surname    string `json:"surname" example:"ivanov"`
	Patronymic string `json:"patronymic" example:"ivanovich"`
}
