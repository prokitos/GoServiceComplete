package app

import (
	postgres "modular/internal/models"
	services "modular/internal/services/myService"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// функция которая принимает запросы
func MainServer() {

	http.HandleFunc("/delete", deleteGetRequest)
	http.HandleFunc("/insert", insertGetRequest)
	http.HandleFunc("/show", showsSpecGetRequest)
	http.HandleFunc("/showall", showAllGetRequest)
	http.HandleFunc("/update", updateGetRequest)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Error("the server cannot start")
		log.Debug("the server cannot use this port")
	}
}

// ответ на запрос для изменения пользователя
func updateGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("receiving a update request")

	var updatePerson postgres.User
	updatePerson.Id = 0
	updatePerson.Age = 0

	param3 := r.FormValue("id")
	if param3 == "" {
		param3 = "-1"
	}
	ids, err := strconv.Atoi(param3)
	if err != nil {
		log.Error("the update request was not executed")
		log.Debug("id couldn't convert to a number: " + param3)
		return
	}
	param4 := r.FormValue("age")
	if param4 == "" {
		param4 = "-1"
	}
	ages, err := strconv.Atoi(param4)
	if err != nil {
		log.Error("the update request was not executed")
		log.Debug("age couldn't convert to a number: " + param4)
		return
	}
	updatePerson.Id = ids
	updatePerson.Age = ages
	updatePerson.Name = r.FormValue("name")
	updatePerson.Surname = r.FormValue("surname")
	updatePerson.Patronymic = r.FormValue("patronymic")
	updatePerson.Nationality = r.FormValue("nationality")
	updatePerson.Sex = r.FormValue("sex")

	services.UpdateData(w, updatePerson)
}

// ответ на запрос для просмотра пользователя по параметрам
func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("receiving a show request")

	var sortPerson string = ""
	var searchPerson postgres.User
	searchPerson.Id = 0
	searchPerson.Age = 0

	// данные для пагинации
	param1 := r.FormValue("limit")
	if param1 == "" {
		param1 = "0"
	}
	limit, err := strconv.Atoi(param1)
	if err != nil {
		log.Error("the show request was not executed")
		log.Debug("limit params couldn't convert to a number: " + param1)
		return
	}
	param2 := r.FormValue("offset")
	if param2 == "" {
		param2 = "0"
	}
	offset, err2 := strconv.Atoi(param2)
	if err2 != nil {
		log.Error("the show request was not executed")
		log.Debug("offset params couldn't convert to a number: " + param2)
		return
	}

	// данные для вывода пользователя с определенным условием
	param3 := r.FormValue("id")
	if param3 == "" {
		param3 = "-1"
	}
	ids, err := strconv.Atoi(param3)
	if err != nil {
		log.Error("the show request was not executed")
		log.Debug("id couldn't convert to a number: " + param3)
		return
	}

	param4 := r.FormValue("age")
	if param4 == "" {
		param4 = "-1"
	}
	ages, err := strconv.Atoi(param4)
	if err != nil {
		log.Error("the show request was not executed")
		log.Debug("age couldn't convert to a number: " + param4)
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

	services.ShowSpecData(w, offset, limit, sortPerson, searchPerson)
}

// ответ на запрос для просмотра всех пользователей
func showAllGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("receiving a show request")
	services.ShowAllData(w)
}

// ответ на запрос для удаления пользователя по айди
func deleteGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("receiving a delete request")
	param1 := r.FormValue("id")
	inter, err := strconv.Atoi(param1)
	if err != nil {
		log.Error("the delete request was not executed")
		log.Debug("id couldn't convert to a number: " + param1)
		return
	}

	services.DeleteDataEncrichment(w, inter)
}

// ответ на запрос для создания нового пользователя
func insertGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("receiving a create request")

	param1 := r.FormValue("name")
	param2 := r.FormValue("surname")
	param3 := r.FormValue("patronymic")

	if len(param1) > 40 || len(param2) > 40 || len(param3) > 40 || len(param1) == 0 {
		log.Error("the insert request was not executed")
		log.Debug("incorrect length of persons data: " + "name=" + param1 + " surname=" + param2 + "patronymic=" + param3)
		return
	}

	var newPerson postgres.User
	newPerson.Name = param1
	newPerson.Surname = param2
	newPerson.Patronymic = param3

	services.CreateDataEncrichment(w, newPerson)
}
