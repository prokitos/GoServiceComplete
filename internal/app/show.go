package app

import (
	"modular/internal/models"
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Users godoc
// @Summary Show persons in users
// @Description Show persons in users
// @Tags users
// @Accept  json
// @Produce  json
// @Param sort query string false "Sort records"
// @Param limit query int false "Show max limit records"
// @Param offset query int false "Show records with current offset"
// @Param user query showUser false "Update user"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Success 200 {object} showUser "successful operation"
// @Router /show [get]
func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a show request")

	var user models.User
	user.Id = r.FormValue("id")
	user.Age = r.FormValue("age")
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Patronymic = r.FormValue("patronymic")
	user.Nationality = r.FormValue("nationality")
	user.Sex = r.FormValue("gender")

	var sortPerson string = r.FormValue("sort")
	var offset string = r.FormValue("offset")
	var limit string = r.FormValue("limit")

	if _, err := strconv.Atoi(offset); offset != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("offset params couldn't convert to a number, offset = " + offset)

		models.BadResponseSend(w, "operation failed, wrong offset = "+offset, 400)
		return
	}
	if _, err := strconv.Atoi(limit); limit != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("limit params couldn't convert to a number, limit = " + limit)

		models.BadResponseSend(w, "operation failed, wrong limit = "+limit, 400)
		return
	}
	if _, err := strconv.Atoi(user.Id); user.Id != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("id params couldn't convert to a number, id = " + user.Id)

		models.BadResponseSend(w, "operation failed, wrong id = "+user.Id, 400)
		return
	}
	if _, err := strconv.Atoi(user.Age); user.Age != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("age params couldn't convert to a number, age = " + user.Age)

		models.BadResponseSend(w, "operation failed, wrong age = "+user.Age, 400)
		return
	}

	services.ShowSpecData(w, offset, limit, sortPerson, user)
}

type showUser struct {
	ID          int    `json:"id" example:"1" format:"int64"`
	Age         int    `json:"age" example:"25" format:"int64"`
	Name        string `json:"name" example:"ivan"`
	Surname     string `json:"surname" example:"ivanov"`
	Patronymic  string `json:"patronymic" example:"ivanovich"`
	Sex         string `json:"gender" example:"male"`
	Nationality string `json:"nationality" example:"RU"`
}
