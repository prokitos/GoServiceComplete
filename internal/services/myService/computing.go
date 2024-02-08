package services

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// вывод возраста из json файла
func AgeComputing(jsonString string) string {

	var instanse messageAge
	data := []byte(jsonString)

	err := json.Unmarshal(data, &instanse)
	if err != nil {
		log.Error("error in getting the age")
		log.Debug("error when getting the age in the json file: " + jsonString)
	}

	return strconv.Itoa(instanse.Age)
}

// вывод пола из json файла
func SexComputing(jsonString string) string {

	var instanse messageSex
	data := []byte(jsonString)

	err := json.Unmarshal(data, &instanse)
	if err != nil {
		log.Error("error in getting the gender")
		log.Debug("error when getting the gender in the json file: " + jsonString)
	}

	return instanse.Gender
}

// вывод наиболее вероятной национальности из json файла
func NationComputing(jsonString string) string {

	var instanse messageNational
	data := []byte(jsonString)

	err := json.Unmarshal(data, &instanse)
	if err != nil {
		log.Error("error in getting the nationality")
		log.Debug("error when getting the nationality in the json file: " + jsonString)
	}

	if len(instanse.Country) == 0 {
		return "NO"
	} else {
		return instanse.Country[0].Country_id
	}

}

type messageAge struct {
	Count int
	Name  string
	Age   int
}

type messageSex struct {
	Count  int
	Name   string
	Gender string
}

type messageNational struct {
	Count   int
	Name    string
	Country []messageCounty
}

type messageCounty struct {
	Country_id  string
	Probability float32
}
