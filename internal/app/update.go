package app

import (
	postgres "modular/internal/models"
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
// @Param id query int true "Update user"
// @Param user body updateUser true "Update user"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /update [put]
func updateGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a update request")

	var updatePerson postgres.User

	param3 := r.FormValue("id")
	param4 := r.FormValue("age")
	if param4 == "" {
		param4 = "-1"
	}

	if param3 == "" {
		log.Error("the update request was not executed")
		log.Debug("id couldn't writing " + param3)
		w.Write([]byte(`"message": "Update request failed"`))
		return
	}
	ids, err := strconv.Atoi(param3)
	if err != nil {
		log.Error("the update request was not executed")
		log.Debug("id couldn't convert to a number: " + param3)
		w.Write([]byte(`"message": "Update request failed"`))
		return
	}
	ages, err := strconv.Atoi(param4)
	if err != nil {
		log.Error("the update request was not executed")
		log.Debug("age couldn't convert to a number: " + param4)
		w.Write([]byte(`"message": "Update request failed"`))
		return
	}

	updatePerson.Id = ids
	updatePerson.Age = ages
	updatePerson.Name = r.FormValue("name")
	updatePerson.Surname = r.FormValue("surname")
	updatePerson.Patronymic = r.FormValue("patronymic")
	updatePerson.Nationality = r.FormValue("nationality")
	updatePerson.Sex = r.FormValue("sex")

	if len(updatePerson.Name) > 40 || len(updatePerson.Surname) > 40 || len(updatePerson.Patronymic) > 40 || len(updatePerson.Sex) > 40 || updatePerson.Age > 200 || len(updatePerson.Nationality) > 40 {
		log.Error("the update request was not executed")
		log.Debug("incorrect length of persons data: ")
		w.Write([]byte(`"message": "Update request failed"`))
		return
	}

	w.Write([]byte(`"message": "Update request succes",`))
	services.UpdateData(w, updatePerson)
}

type updateUser struct {
	Name       string `json:"name" example:"ivan"`
	Surname    string `json:"surname" example:"ivanov"`
	Patronymic string `json:"patronymic" example:"ivanovich"`
}
