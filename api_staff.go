package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Структуры для API
type StaffInfo struct {
	Id                int64  `json:"id"`
	Fio               string `json:"fio"`
	Phone             string
	Position          string
	Signer            int64 `json:"signer"`
	Fio_Dative        string
	Position_Dative   string
	Fio_Genitive      string
	Position_Genitive string
}

func ApiStaffList(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	res = dbkeeper.StaffList()
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("ApiStaffList: error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ApiStaffList: write error %v", err)
		return
	}
}

//API
func (dk *DBKeeper) StaffList() DBResult {
	var resultData []StaffInfo = make([]StaffInfo, 0)
	err := dk.db.Select(&resultData, "select id, fio, signer from staff order by fio;")
	if err != nil {
		log.Printf("DBKeeper.StaffList: select query error: %v\n", err)
		return DBResult{err, nil}
	}
	return DBResult{nil, resultData}
}
