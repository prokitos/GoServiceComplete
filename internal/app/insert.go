package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	param1 := r.FormValue("name")
	param2 := r.FormValue("surname")
	param3 := r.FormValue("patronymic")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post addUser
	json.Unmarshal(reqBody, &post)
	fmt.Print(post.Name)

	if len(param1) > 40 || len(param2) > 40 || len(param3) > 40 || len(param1) == 0 || len(param2) == 0 || len(param3) == 0 {
		log.Error("the insert request was not executed")
		log.Debug("incorrect length of persons data: " + "name=" + param1 + " surname=" + param2 + "patronymic=" + param3)
		w.Write([]byte(`"message": "Insert request failed"`))
		return
	}

	var newPerson postgres.User
	newPerson.Name = param1
	newPerson.Surname = param2
	newPerson.Patronymic = param3

	w.Write([]byte(`"message": "Insert request succes",`))
	services.CreateDataEncrichment(w, newPerson)
}

type addUser struct {
	Name       string `json:"name" example:"ivan"`
	Surname    string `json:"surname" example:"ivanov"`
	Patronymic string `json:"patronymic" example:"ivanovich"`
}
