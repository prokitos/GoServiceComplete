package models

// структура таблицы
type User struct {
	Id          string
	Name        string
	Surname     string
	Patronymic  string
	Age         string
	Sex         string
	Nationality string
}

type UserTemp struct {
	Id          int
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Sex         string
	Nationality string
}
