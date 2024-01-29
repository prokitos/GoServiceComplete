package postgres

// структура таблицы
type User struct {
	Id          int
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Sex         string
	Nationality string
}
