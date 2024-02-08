package services

import (
	"io/ioutil"
	table_user "modular/internal/database/db_postgress"
	postgres "modular/internal/models"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// Обновление записи
func UpdateData(w http.ResponseWriter, user postgres.User) {
	log.Info("trying to update records with id " + user.Id)

	var conn string = table_user.ConStringUpdate(user)
	dbCon := table_user.ConnectToDb("internal/config/postgress.env")
	table_user.ExecuteToDB(dbCon, w, conn, "Update")
}

// Удаление записи по айди
func DeleteDataEncrichment(w http.ResponseWriter, id string) {
	log.Info("trying to delete records with id " + id)

	var conn string = table_user.ConStringDelete(id)
	dbCon := table_user.ConnectToDb("internal/config/postgress.env")
	table_user.ExecuteToDB(dbCon, w, conn, "Delete")
}

// Показать записи
func ShowSpecData(w http.ResponseWriter, offset string, limit string, sort string, user postgres.User) {
	log.Info("trying to showing data")
	var conn string = table_user.ConStringShowSpec(offset, limit, sort, user)
	dbCon := table_user.ConnectToDb("internal/config/postgress.env")
	table_user.ShowFromDB(dbCon, w, conn)
}

// Создать новую запись
func CreateDataEncrichment(w http.ResponseWriter, user postgres.User) {
	log.Info("trying to create data with name=" + user.Name + " surname=" + user.Surname + " patronymic=" + user.Patronymic)

	chann1 := make(chan string)
	chann2 := make(chan string)
	chann3 := make(chan string)

	go getAgeFromName(user.Name, chann1)
	go getSexFromName(user.Name, chann2)
	go getNationalityFromName(user.Name, chann3)

	user.Age = <-chann1
	user.Sex = <-chann2
	user.Nationality = <-chann3

	log.Info("A response was received from all the APIs")

	var conn string = table_user.ConStringInsert(user)

	dbCon := table_user.ConnectToDb("internal/config/postgress.env")
	table_user.ExecuteToDB(dbCon, w, conn, "Create")
}

// получить национальность по имени
func getNationalityFromName(p_name string, chans chan string) {
	var result string = sendRequestToGet("https://api.nationalize.io/", p_name)
	chans <- NationComputing(result)
}

// получить пол по имени
func getSexFromName(p_name string, chans chan string) {
	var result string = sendRequestToGet("https://api.genderize.io/", p_name)
	chans <- SexComputing(result)
}

// получить возраст по имени
func getAgeFromName(p_name string, chans chan string) {
	var result string = sendRequestToGet("https://api.agify.io/", p_name)
	chans <- AgeComputing(result)
}

// отправка гет запроса на указанный сайт с указанным параметром
func sendRequestToGet(curUrl string, p_name string) string {

	baseURL, _ := url.Parse(curUrl)
	params := url.Values{}
	params.Add("name", p_name)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Debug("Error connecting to external api")
		log.Error("Error adding a user")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var results string = string(body)
	return results
}
