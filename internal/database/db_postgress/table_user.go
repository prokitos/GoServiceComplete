package db_postgress

import (
	"database/sql"
	"encoding/json"
	"fmt"
	postgres "modular/internal/models"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// функция которая вызывает миграцию через код
// func migrateStart() {
// 	cmd := exec.Command("goose", "-dir", "db/migrations", "postgres", "postgresql://postgres:root@127.0.0.1:8092/postgres?sslmode=disable", "up")
// 	err := cmd.Run()
// 	if err != nil {
// 		println("panic !!!")
// 	}
// 	println("Migrate complete")
// }

func ConnectToDb(path string) *sql.DB {

	log.Info("connecting to the database")

	godotenv.Load(path)

	envUser := os.Getenv("USER")
	envPass := os.Getenv("PASS")
	envHost := os.Getenv("HOST")
	envPort := os.Getenv("PORT")
	envName := os.Getenv("NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
	}

	db.Begin()

	return db
}

// вызов операции над таблицей
func ExecuteToDB(db *sql.DB, w http.ResponseWriter, conn string, operation string) {
	defer db.Close()

	result, err := db.Exec(conn)
	if err != nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + conn)
	}
	if result == nil {
		w.Write([]byte(`"status": "400"`))
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("database connection error")
		log.Debug("data change error: " + conn)
	}

	// Логи и вывод данных на сервер/клиент
	log.Info(operation + "complete")
	fmt.Print("operation " + operation + ", ")
	fmt.Print(rowsAffected)
	fmt.Println(" Rows affected")

	var affectedRow int = int(rowsAffected)
	errResp := postgres.GoodResponse{
		Message:  "success operation",
		Code:     200,
		Affected: affectedRow,
	}
	json.NewEncoder(w).Encode(errResp)
}

// показать таблицу
func ShowFromDB(db *sql.DB, w http.ResponseWriter, stroka string) {

	defer db.Close()

	rows, err := db.Query(stroka)
	if err != nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + stroka)
	}
	if rows == nil {
		w.Write([]byte(`"status": "Null execute"`))
		return
	}
	defer rows.Close()
	users := []postgres.User{}

	for rows.Next() {
		p := postgres.User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Sex, &p.Nationality)
		if err != nil {
			log.Error("database show data error")
			log.Debug("errored query when show data: " + stroka)
			continue
		}
		users = append(users, p)
	}

	w.Write([]byte(`{`))
	for _, i := range users {
		w.Write([]byte(`{"id": "` + strconv.Itoa(i.Id) + `",`))
		w.Write([]byte(`"name": "` + i.Name + `",`))
		w.Write([]byte(`"surname": "` + i.Surname + `",`))
		w.Write([]byte(`"patronymic": "` + i.Patronymic + `",`))
		w.Write([]byte(`"age": "` + strconv.Itoa(i.Age) + `",`))
		w.Write([]byte(`"gender": "` + i.Sex + `",`))
		w.Write([]byte(`"nationality": "` + i.Nationality + `"},`))
	}
	w.Write([]byte(`}`))

	log.Info("the data was successfully shown to the user")
}
