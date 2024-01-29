package db_postgress

import (
	"database/sql"
	"fmt"
	postgres "modular/internal/models"
	"net/http"
	"os"

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

// вызов операции над таблицей
func ExecuteToDB(w http.ResponseWriter, conn string, operation string) {
	log.Info("connecting to the database")

	godotenv.Load()
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
	defer db.Close()

	result, err := db.Exec(conn)
	if err != nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + conn)
	}

	// Логи и вывод данных на сервер/клиент
	log.Info(operation + " complete")
	fmt.Print(operation)
	fmt.Print(result.RowsAffected())
	fmt.Println(" Rows affected")
	fmt.Fprintln(w, "data has", operation)
}

// показать таблицу
func ShowFromDB(w http.ResponseWriter, stroka string) {
	log.Info("connecting to the database")

	godotenv.Load()
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
	defer db.Close()

	rows, err := db.Query(stroka)
	if err != nil {
		log.Error("database connection error")
		log.Debug("database error executing the request: " + stroka)
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

	fmt.Fprintln(w, " ")
	for _, i := range users {
		fmt.Fprintln(w, i.Id, i.Name, i.Surname, i.Patronymic, i.Age, i.Sex, i.Nationality)
	}

	log.Info("the data was successfully shown to the user")
}
