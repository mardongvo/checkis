package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

//Главная структура
//Поддерживает содинение с БД, выполняет запросы
type DBKeeper struct {
	database string //database connection string
	db       *sqlx.DB
}

//Результат запроса
//Требется диагностировать ошибку, поэтому так обернуто
type DBResult struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}

func NewDBKeeper(database string) DBKeeper {
	var err error
	dk := DBKeeper{database, nil}
	dk.db, err = sqlx.Open("postgres", dk.database)
	if err != nil {
		log.Printf("DBKeeper: подключение к БД не удалось, ошибка %v", err)
	}
	return dk
}

func (dk *DBKeeper) Close() {
	if dk.db != nil {
		dk.db.Close()
	}
}
