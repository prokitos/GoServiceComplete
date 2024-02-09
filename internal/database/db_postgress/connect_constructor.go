package db_postgress

import (
	postgres "modular/internal/models"
	"reflect"
	"strconv"
	"strings"
)

// создание строки запроса для обновления записи
func ConStringUpdate(user postgres.User) string {

	var stroka string = "update users "
	var additional string = ""

	if user.Name != "" {
		additional += " name = '" + user.Name + "' ,"
	}
	if user.Surname != "" {
		additional += " surname = '" + user.Surname + "' ,"
	}
	if user.Patronymic != "" {
		additional += " patronymic = '" + user.Patronymic + "' ,"
	}
	if user.Sex != "" {
		additional += " sex = '" + user.Sex + "' ,"
	}
	if user.Nationality != "" {
		additional += " nationality = '" + user.Nationality + "' ,"
	}
	if user.Age != "-1" {
		additional += " age = '" + user.Age + "' ,"
	}

	additional = strings.TrimSuffix(additional, ",")
	if len(additional) > 0 {
		stroka += "set" + additional
	}

	stroka += "where id = '" + user.Id + "'"

	return stroka
}

// создание строки запроса для просмотра записей
func ConStringShowSpec(offset string, limit string, sort string, user postgres.User) string {
	// создание строки по параметрам
	var stroka string = "select * from users"
	var whereCheck bool = false

	values := reflect.ValueOf(user)
	for i := 0; i < values.NumField(); i++ {

		var curName string = values.Type().Field(i).Name
		var curValue string = ""

		// если пришло число, то конвертить из инта в стринг, иначе просто брать стринг
		if values.Field(i).Type().Name() == "int" {
			temp := values.Field(i).Interface().(int)
			if temp >= 0 {
				curValue = strconv.Itoa(temp)
			}
		} else {
			curValue = values.Field(i).String()
		}

		if curValue != "" {
			if !whereCheck {
				stroka += " where " + (strings.ToLower(curName)) + " = '" + curValue + "'"
				whereCheck = true
			} else {
				stroka += " and " + (strings.ToLower(curName)) + " = '" + curValue + "'"
			}
		}

	}

	if len(sort) > 0 {
		stroka += " order by " + sort
	}
	if limit != "" {
		stroka += " limit " + limit
	}
	if offset != "" {
		stroka += " offset " + offset
	}

	return stroka
}

// создание строки запроса для удаления записи по айди
func ConStringDelete(id string) string {
	var stroka string = "delete from users where id = '" + id + "'"
	return stroka
}

// создание строки для создания записи
func ConStringInsert(user postgres.User) string {
	var result string = "insert into users (name, surname, patronymic, age, sex, nationality) values ('" + user.Name + "', '" +
		user.Surname + "', '" + user.Patronymic + "', '" + user.Age + "', '" + user.Sex + "', '" + user.Nationality + "')"
	return result
}
