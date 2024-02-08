package app

import (
	"encoding/json"
	"io"
	"modular/internal/models"
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Users godoc
// @Summary Update persons in users
// @Description Update persons in users
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body updateUser true "Update user"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /update [put]
func updateGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a update request")

	reqBody, _ := io.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(reqBody, &user)

	if _, err := strconv.Atoi(user.Id); err != nil {
		models.BadResponseSend(w, "operation failed, wrong id = "+user.Id, 400)
		return
	}

	if _, err := strconv.Atoi(user.Age); user.Age != "" && err != nil {
		models.BadResponseSend(w, "operation failed, wrong Age = "+user.Age, 400)
		return
	}

	if len(user.Name) > 40 || len(user.Surname) > 40 || len(user.Patronymic) > 40 || len(user.Sex) > 40 || len(user.Nationality) > 40 {
		log.Error("the update request was not executed")
		log.Debug("incorrect length of persons data: ")

		models.BadResponseSend(w, "operation failed, wrong data lenght", 400)
		return
	}

	services.UpdateData(w, user)
}

type updateUser struct {
	Id         string `json:"id" example:"5"`
	Name       string `json:"name" example:"ivan"`
	Surname    string `json:"surname" example:"ivanov"`
	Patronymic string `json:"patronymic" example:"ivanovich"`
}
