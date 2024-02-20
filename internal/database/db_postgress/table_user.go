package db_postgress

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"modular/internal/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

func ConnectToDb(path string) *sql.DB {

	log.Info("connecting to the database")

	godotenv.Load(path)

	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
	}

	db.Begin()

	return db
}

func MigrateStart() {

	duration := time.Second * 5
	time.Sleep(duration)

	db := ConnectToDb("internal/config/postgress.env")

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "internal/database/migrations"); err != nil {
		log.Error("migration connection error")
		panic(err)
	}
}

// вызов операции над таблицей
func ExecuteToDB(db *sql.DB, w http.ResponseWriter, conn string, operation string) {
	defer db.Close()

	result, err := db.Exec(conn)
	if err != nil || result == nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + conn)

		models.BadResponseSend(w, "operation failed, does not connect to database", 400)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("database connection error")
		log.Debug("data change error: " + conn)

		models.BadResponseSend(w, "operation failed, non valide query", 400)
		return
	}

	// Логи и вывод данных на сервер/клиент
	if rowsAffected > 0 {
		var affectedRow int = int(rowsAffected)
		log.Info(operation + " complete")
		log.Info(strconv.Itoa(affectedRow) + " Rows affected")

		models.GoodResponseSend(w, "success operation", affectedRow)
	} else {
		log.Info(operation + " not complete, records does not exist!!")

		models.BadResponseSend(w, "operation failed, nothing to execute", 404)
	}

}

// показать таблицу
func ShowFromDB(db *sql.DB, w http.ResponseWriter, stroka string) {

	defer db.Close()

	rows, err := db.Query(stroka)
	if err != nil || rows == nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + stroka)

		models.BadResponseSend(w, "operation failed, does not connect to database", 400)
		return
	}
	defer rows.Close()
	users := []models.User{}

	for rows.Next() {
		p := models.User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Sex, &p.Nationality)
		if err != nil {
			log.Error("database show data error")
			log.Debug("errored query when show data: " + stroka)

			models.BadResponseSend(w, "operation failed, does not connect to database", 400)
			continue
		}
		users = append(users, p)
	}

	mass := []models.User{}
	for _, i := range users {
		mass = append(mass, i)
	}

	if len(mass) == 0 {
		models.BadResponseSend(w, "does not records to show", 404)
		log.Info("nothing to show")
		return
	}

	json.NewEncoder(w).Encode(mass)

	log.Info("the data was successfully shown to the user")
}
