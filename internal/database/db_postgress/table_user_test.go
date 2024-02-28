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
			conn:    "insert into test (name, surname, patronymic, age, sex, nationality) values ('denis', 'denisov', 'denisovich', '50', 'male', 'RU')",
			oper:    "Insert",
			want:    `{"message":"success operation","code":200,"affected_rows":1} insert` + "\n",
		},
		{
			testNum: 2,
			name:    "test update",
			conn:    "update test set sex = 'female' , nationality = 'RU' , age = '50' where id = '2'",
			oper:    "Update",
			want:    `{"message":"success operation","code":200,"affected_rows":1} update` + "\n",
		},
		{
			testNum: 3,
			name:    "test delete",
			conn:    "delete from test where id = '99'",
			oper:    "Delete",
			want:    `{"message":"operation failed, nothing to execute","code":404}` + "\n",
		},
	}

	for _, tc := range testTable {

		connStr := GetConnectString("../../config/postgress.env")

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
			conn:    "select * from test where id = '1'",
			want:    `[{"Id":"1","Name":"denis","Surname":"denisov","Patronymic":"denisovich","Age":"50","Sex":"male","Nationality":"RU"}]` + "\n",
		},
		{
			testNum: 2,
			name:    "offset 2 limit 1",
			conn:    "select * from test order by id offset 2 limit 1",
			want:    `[{"Id":"3","Name":"denis","Surname":"denisov","Patronymic":"denisovich","Age":"50","Sex":"male","Nationality":"RU"}]` + "\n",
		},
	}

	for _, tc := range testTable {

		connStr := GetConnectString("../../config/postgress.env")
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

	var res *sql.DB = ConnectToDb("../../config/postgress.env")

	if res.Stats().InUse == 0 {
		t.Errorf("result wrong at test, does not connect to server")
	}

	res.Close()

}
