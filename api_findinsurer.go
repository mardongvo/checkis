package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Структуры для API
type InsurerInfo struct {
	Regnum      string `json:"insurer_regnum"`
	Kpsnum      string `json:"insurer_kpsnum"`
	INN         string `json:"insurer_inn"`
	KPP         string `json:"insurer_kpp"`
	Fullname    string `json:"insurer_fullname"`
	Shortname   string `json:"insurer_shortname"`
	Postindex   string `json:"insurer_postindex"`
	Postaddress string `json:"insurer_postaddress"`
}

func ApiFindInsurer(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	vars := mux.Vars(r)
	regnum := vars["regnum"]
	if len(regnum) < 1 {
		res = DBResult{fmt.Errorf("ApiFindInsurer: Длина строки поиска должна быть не меньше 1 символа"), nil}
	} else {
		res = dbkeeper.FindInsurer(regnum)
	}
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("ApiFindInsurer: error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ApiFindInsurer: write error %v", err)
		return
	}
}

//API
func (dk *DBKeeper) FindInsurer(regnum string) DBResult {
	if regnum == "" {
		return DBResult{fmt.Errorf("Строка поиска пуста"), nil}
	}
	var resultError error = nil
	var resultData InsurerInfo
	err := dk.db.Get(&resultData, "select * from insurers where regnum=$1;", regnum)
	if err != nil {
		log.Printf("DBKeeper.FindInsurer: select query error: %v\n", err)
		return DBResult{err, nil}
	}
	return DBResult{resultError, resultData}
}
