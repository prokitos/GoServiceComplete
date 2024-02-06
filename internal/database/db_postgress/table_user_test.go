package db_postgress

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestExecuteDB(t *testing.T) {

	testTable := []struct {
		testNum int
		name    string
		conn    string
		oper    string
		want    string
	}{
		{
			testNum: 1,
			name:    "test insert",
			conn:    "insert into testuser (name, surname, patronymic, age, sex, nationality) values ('denis', 'denisov', 'denisovich', '50', 'male', 'RU')",
			oper:    "Insert",
			want:    "\"status\": \"Insert success\",\"affected_rows\": \"1\"",
		},
		{
			testNum: 2,
			name:    "test update",
			conn:    "update testuser set sex = 'female' , nationality = 'RU' , age = '50' where id = '2'",
			oper:    "Update",
			want:    "\"status\": \"Update success\",\"affected_rows\": \"1\"",
		},
		{
			testNum: 3,
			name:    "test delete",
			conn:    "delete from testuser where id = '99'",
			oper:    "Delete",
			want:    "\"status\": \"Delete success\",\"affected_rows\": \"0\"",
		},
	}

	for _, tc := range testTable {

		connStr := "postgresql://postgres:root@127.0.0.1:8092/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Error("database connection error")
			log.Debug("there is not connection with database")
		}

		rr := httptest.NewRecorder()
		ExecuteToDB(db, rr, tc.conn, tc.oper)
		if rr.Body.String() != tc.want {
			t.Errorf("result wrong at test, got [%v] want [%v]", rr.Body, tc.want)
		}

	}
}

func TestShowDB(t *testing.T) {

	testTable := []struct {
		testNum int
		name    string
		conn    string
		want    string
	}{
		{
			testNum: 1,
			name:    "show id = 1",
			conn:    "select * from testuser where id = '1'",
			want:    "{{\"id\": \"1\",\"name\": \"denis\",\"surname\": \"denisov\",\"patronymic\": \"denisovich\",\"age\": \"50\",\"gender\": \"male\",\"nationality\": \"RU\"},}",
		},
		{
			testNum: 2,
			name:    "offset 2 limit 1",
			conn:    "select * from testuser order by id offset 2 limit 1",
			want:    "{{\"id\": \"3\",\"name\": \"test\",\"surname\": \"denisov\",\"patronymic\": \"denisovich\",\"age\": \"50\",\"gender\": \"male\",\"nationality\": \"RU\"},}",
		},
	}

	for _, tc := range testTable {

		connStr := "postgresql://postgres:root@127.0.0.1:8092/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Error("database connection error")
			log.Debug("there is not connection with database")
		}

		rr := httptest.NewRecorder()
		ShowFromDB(db, rr, tc.conn)
		if rr.Body.String() != tc.want {
			t.Errorf("result wrong at test, got [%v] want [%v]", rr.Body, tc.want)
		}

	}
}

func TestConnetion(t *testing.T) {

	res := ConnectToDb("../../config/postgress.env")

	if res.Stats().InUse == 0 {
		t.Errorf("result wrong at test, does not connect to server")
	}

	res.Close()

}
