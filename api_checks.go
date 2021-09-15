package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const COMMON_DATE_LAYOUT = "2006-01-02"

func checkFormat(inp string, typ string) bool {
	if typ == "date" {
		_, err := time.Parse(COMMON_DATE_LAYOUT, inp)
		if err != nil {
			return false
		}
		return true
	}
	if typ == "int" {
		_, err := strconv.ParseInt(inp, 10, 64)
		if err != nil {
			return false
		}
		return true
	}
	return true
}

func ApiListChecks(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	var err error
	var filter TFilter
	filter.regnum = strings.Trim(r.URL.Query().Get("regnum"), " \t")
	if checkFormat(r.URL.Query().Get("date1"), "date") {
		filter.date1 = r.URL.Query().Get("date1")
	}
	if checkFormat(r.URL.Query().Get("date2"), "date") {
		filter.date2 = r.URL.Query().Get("date2")
	}
	if checkFormat(r.URL.Query().Get("inspector"), "int") {
		filter.inspector, _ = strconv.ParseInt(r.URL.Query().Get("inspector"), 10, 64)
	}
	res = dbkeeper.SelectChecks(filter)
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("ApiFilterChecks: error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ApiFilterChecks: write error %v", err)
		return
	}
}

func ApiGetCheck(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	var id int64
	var err error
	id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		res = DBResult{fmt.Errorf("ApiGetCheck: ошибка разбора id: %v", err), nil}
		goto FALL
	}
	res = dbkeeper.GetCheckById(id, true)
FALL:
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("ApiGetCheck: error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ApiGetCheck: write error %v", err)
		return
	}
}

func ApiGetPart(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	var id int64
	var err error
	vars := mux.Vars(r)
	part := vars["part"]
	_id := vars["id"]
	id, err = strconv.ParseInt(_id, 10, 64)
	if err != nil {
		res = DBResult{fmt.Errorf("ApiGetCheck: ошибка разбора id: %v", err), nil}
		goto FALL
	}
	switch part {
	case "actlist":
		res = dbkeeper.GetActList(id)
	}
FALL:
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("ApiGetCheck: error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ApiGetCheck: write error %v", err)
		return
	}
}

func ApiSaveCheck(w http.ResponseWriter, r *http.Request) {
	var res DBResult
	var rq map[string]interface{} = make(map[string]interface{})
	part := mux.Vars(r)["part"]

	inp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("JsonApiSaveCheck(1): error %v", err)
		res = DBResult{err, nil}
		goto END
	}
	err = json.Unmarshal(inp, &rq)
	if err != nil {
		log.Printf("JsonApiSaveCheck(2): error %v", err)
		res = DBResult{err, nil}
		goto END
	}
	switch part {
	case "header":
		res = dbkeeper.UpsertCheck(rq, CMD_NONE)
	case "req":
		res = dbkeeper.UpsertCheck(rq, CMD_DOCS_REQ)
	case "memorandum":
		res = dbkeeper.UpsertCheck(rq, CMD_DOCS_MEMORANDUM)
	case "act":
		res = dbkeeper.UpsertCheck(rq, CMD_DOCS_ACT)
	case "decision":
		res = dbkeeper.UpsertCheck(rq, CMD_DOCS_DECISION)
	case "charge":
		res = dbkeeper.UpsertCheck(rq, CMD_DOCS_CHARGE)
	case "actlist":
		res = dbkeeper.UpsertActList(rq)
	}
END:
	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("JsonApiSaveCheck(3): error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("JsonApiSaveCheck: write error %v", err)
		return
	}
}
