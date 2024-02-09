package db_postgress

import (
	postgres "modular/internal/models"
	"testing"
)

func TestInsertConstructor(t *testing.T) {

	testTable := []struct {
		testNum int
		char    postgres.User
		name    string
		want    string
	}{
		{
			testNum: 1,
			name:    "normal query",
			char:    postgres.User{Id: "6", Name: "denis", Surname: "denisov", Patronymic: "denisovich", Age: "50", Sex: "male", Nationality: "RU"},
			want:    "insert into users (name, surname, patronymic, age, sex, nationality) values ('denis', 'denisov', 'denisovich', '50', 'male', 'RU')",
		},
		{
			testNum: 2,
			name:    "without id",
			char:    postgres.User{Name: "denis", Surname: "denisov", Patronymic: "denisovich", Age: "50", Sex: "male", Nationality: "RU"},
			want:    "insert into users (name, surname, patronymic, age, sex, nationality) values ('denis', 'denisov', 'denisovich', '50', 'male', 'RU')",
		},
	}

	for _, tc := range testTable {

		var newPerson postgres.User = tc.char
		result := ConStringInsert(newPerson)

		// Обработчиик ошибок
		if result != tc.want {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, result, tc.want)
		}
	}
}

func TestDeleteConstructor(t *testing.T) {

	testTable := []struct {
		testNum int
		id      string
		name    string
		want    string
	}{
		{
			testNum: 1,
			name:    "normal query",
			id:      "6",
			want:    "delete from users where id = '6'",
		},
		{
			testNum: 2,
			name:    "normal query 2",
			id:      "0",
			want:    "delete from users where id = '0'",
		},
	}

	for _, tc := range testTable {

		result := ConStringDelete(tc.id)

		// Обработчиик ошибок
		if result != tc.want {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, result, tc.want)
		}
	}
}

func TestUpdateConstructor(t *testing.T) {

	testTable := []struct {
		testNum int
		char    postgres.User
		name    string
		want    string
	}{
		{
			testNum: 1,
			name:    "normal query",
			char:    postgres.User{Id: "6", Name: "denis", Surname: "denisov", Patronymic: "denisovich", Age: "50", Sex: "male", Nationality: "RU"},
			want:    "update users set name = 'denis' , surname = 'denisov' , patronymic = 'denisovich' , sex = 'male' , nationality = 'RU' , age = '50' where id = '6'",
		},
		{
			testNum: 2,
			name:    "normal query 2",
			char:    postgres.User{Id: "2", Age: "50", Sex: "male", Nationality: "RU"},
			want:    "update users set sex = 'male' , nationality = 'RU' , age = '50' where id = '2'",
		},
	}

	for _, tc := range testTable {

		var newPerson postgres.User = tc.char
		result := ConStringUpdate(newPerson)

		// Обработчиик ошибок
		if result != tc.want {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, result, tc.want)
		}
	}
}

func TestShowConstructor(t *testing.T) {

	testTable := []struct {
		testNum int
		name    string
		char    postgres.User
		sort    string
		offset  string
		limit   string
		want    string
	}{
		{
			testNum: 1,
			name:    "show all",
			char:    postgres.User{Id: "", Age: ""},
			want:    "select * from users",
		},
		{
			testNum: 2,
			name:    "show all by order with limit and offset",
			char:    postgres.User{Id: "", Age: ""},
			offset:  "2",
			limit:   "10",
			sort:    "name",
			want:    "select * from users order by name limit 10 offset 2",
		},
		{
			testNum: 3,
			name:    "normal query",
			char:    postgres.User{Name: "denis", Surname: "denisov", Id: "", Age: ""},
			sort:    "id",
			want:    "select * from users where name = 'denis' and surname = 'denisov' order by id",
		},
	}

	for _, tc := range testTable {

		var newPerson postgres.User = tc.char
		result := ConStringShowSpec(tc.offset, tc.limit, tc.sort, newPerson)

		// Обработчиик ошибок
		if result != tc.want {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, result, tc.want)
		}
	}
}
