package app

import (
	postgres "modular/internal/models"
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
// @Param limit query int false "Show max limit records"
// @Param offset query int false "Show records with current offset"
// @Param user query showUser false "Show user"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Success 200 {object} showUser "successful operation"
// @Router /show [get]
func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Info("receiving a show request")

	var sortPerson string = ""
	var searchPerson postgres.User
	// searchPerson.Id = 0
	// searchPerson.Age = 0

	// данные для пагинации
	param1 := r.FormValue("limit")
	param2 := r.FormValue("offset")
	if param1 == "" {
		param1 = "0"
	}
	if param2 == "" {
		param2 = "0"
	}
	limit, err := strconv.Atoi(param1)
	offset, err2 := strconv.Atoi(param2)
	if err != nil || err2 != nil {
		log.Error("the show request was not executed")
		log.Debug("limit or offset params couldn't convert to a number: limit=" + param1 + " offset=" + param2)
		w.Write([]byte(`"message": "Show request failed"`))
		return
	}

	// данные для вывода пользователя с определенным условием
	param3 := r.FormValue("id")
	param4 := r.FormValue("age")
	if param3 == "" {
		param3 = "-1"
	}
	if param4 == "" {
		param4 = "-1"
	}
	ids, err := strconv.Atoi(param3)
	ages, err2 := strconv.Atoi(param4)
	if err != nil || err2 != nil {
		log.Error("the show request was not executed")
		log.Debug("id or age couldn't convert to a number: id=" + param3 + " age=" + param4)
		w.Write([]byte(`"message": "Show request failed"`))
		return
	}

	searchPerson.Id = ids
	searchPerson.Age = ages
	searchPerson.Name = r.FormValue("name")
	searchPerson.Surname = r.FormValue("surname")
	searchPerson.Patronymic = r.FormValue("patronymic")
	searchPerson.Nationality = r.FormValue("nationality")
	searchPerson.Sex = r.FormValue("sex")
	sortPerson = r.FormValue("sort")

	w.Write([]byte(`"message": "Show request succes",`))
	services.ShowSpecData(w, offset, limit, sortPerson, searchPerson)
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
