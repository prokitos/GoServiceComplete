package db_postgress

import (
	postgres "modular/internal/models"
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
	if user.Age >= 0 {
		additional += " age = '" + strconv.Itoa(user.Id) + "' ,"
	}

	additional = strings.TrimSuffix(additional, ",")

	if len(additional) > 0 {
		stroka += "set" + additional
	}

	stroka += "where id = '" + strconv.Itoa(user.Id) + "'"

	return stroka
}

// создание строки запроса для удаления записи по айди
func ConStringDelete(id int) string {
	var stroka string = "delete from users where id = '" + strconv.Itoa(id) + "'"
	return stroka
}

// создание строки запроса для возвращения записей по параметрам
func ConStringShowSpec(offset int, limit int, sort string, user postgres.User) string {
	// создание строки по параметрам
	var stroka string = "select * from users"
	var whereCheck bool = false

	if user.Id >= 0 {
		if !whereCheck {
			stroka += " where id = '" + strconv.Itoa(user.Id) + "'"
			whereCheck = true
		} else {
			stroka += " and id = '" + strconv.Itoa(user.Id) + "'"
		}
	}
	if user.Age >= 0 {
		if !whereCheck {
			stroka += " where age = '" + strconv.Itoa(user.Age) + "'"
			whereCheck = true
		} else {
			stroka += " and age = '" + strconv.Itoa(user.Age) + "'"
		}
	}
	if user.Name != "" {
		if !whereCheck {
			stroka += " where name = '" + user.Name + "'"
			whereCheck = true
		} else {
			stroka += " and name = '" + user.Name + "'"
		}
	}
	if user.Surname != "" {
		if !whereCheck {
			stroka += " where surname = '" + user.Surname + "'"
			whereCheck = true
		} else {
			stroka += " and surname = '" + user.Surname + "'"
		}
	}
	if user.Patronymic != "" {
		if !whereCheck {
			stroka += " where patronymic = '" + user.Patronymic + "'"
			whereCheck = true
		} else {
			stroka += " and patronymic = '" + user.Patronymic + "'"
		}
	}
	if user.Sex != "" {
		if !whereCheck {
			stroka += " where sex = '" + user.Sex + "'"
			whereCheck = true
		} else {
			stroka += " and sex = '" + user.Sex + "'"
		}
	}
	if user.Nationality != "" {
		if !whereCheck {
			stroka += " where nationality = '" + user.Nationality + "'"
			whereCheck = true
		} else {
			stroka += " and nationality = '" + user.Nationality + "'"
		}
	}

	if len(sort) > 0 {
		stroka += " order by " + sort
	}
	if limit > 0 {
		stroka += " limit " + strconv.Itoa(limit)
	}
	if offset > 0 {
		stroka += " offset " + strconv.Itoa(offset)
	}

	return stroka
}

// создание строки чтобы показать все записи
func ConStringShowAll() string {
	return "select * from users"
}

// создание строки для создания записи
func ConStringInsert(user postgres.User) string {
	var result string = "insert into users (name, surname, patronymic, age, sex, nationality) values ('" + user.Name + "', '" +
		user.Surname + "', '" + user.Patronymic + "', '" + strconv.Itoa(user.Age) + "', '" + user.Sex + "', '" + user.Nationality + "')"
	return result
}
