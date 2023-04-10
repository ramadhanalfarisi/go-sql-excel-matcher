package helper

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataSql struct {
	Id   string
	Date string
}

func Connect() *sql.DB {
	strCon := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", "root", "password", "localhost", 3306, "sql_matcher")
	db, err := sql.Open("mysql", strCon)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func ReadData(db *sql.DB, chanIn <-chan string) []DataSql {
	var output []DataSql
	for item := range chanIn {
		sql := `SELECT id,date FROM articles WHERE id = ?`

		rows, err := db.Query(sql, item)
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		if rows != nil {
			for rows.Next() {
				count++
				var (
					id   string
					date string
				)
				if err := rows.Scan(&id, &date); err != nil {
					log.Fatal(err)
				}
				data := DataSql{Id: id, Date: date}
				output = append(output, data)
				log.Println(count, "Data Executed")
			}
		}
	}

	return output
}
