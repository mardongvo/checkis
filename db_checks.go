package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	CMD_NONE            = 0
	CMD_DOCS_REQ        = 1
	CMD_DOCS_MEMORANDUM = 2
	CMD_DOCS_ACT        = 3
	CMD_DOCS_DECISION   = 4
	CMD_DOCS_CHARGE     = 5
)

var CHECKS_KEY_LIST = []string{
	"id",
	//страхователь
	"insurer_regnum",
	"insurer_kpsnum",
	"insurer_inn",
	"insurer_kpp",
	"insurer_fullname",
	"insurer_shortname",
	"insurer_postindex",
	"insurer_postaddress",
	//периоды
	"check_date_start",
	"check_date_end",
	"check_period_start",
	"check_period_end",
	//проверяющий
	"inspector_id",
	"inspector_fio",
	"inspector_position",
	"inspector_phone",
	//1. требование
	"docs_req_number",
	"docs_req_date",
	"docs_req_flag_sicklist",
	"docs_req_flag_birth",
	"docs_req_flag_1_early",
	"docs_req_flag_1_birth",
	"docs_req_flag_15_child",
	"docs_req_flag_accident",
	"docs_req_flag_vacation",
	"docs_req_flag_4days",
	"docs_req_flag_burial_soc",
	"docs_req_flag_burial_spec",
	//2. докладн записка
	"docs_memorandum_number",
	"docs_memorandum_date",
	//3. акт
	"docs_act_number",
	"docs_act_date",
	//4. решение
	"docs_decision_number",
	"docs_decision_date",
	//5. требование об уплате
	"docs_charge_number",
	"docs_charge_date",
}

var CHECKS_KEY_DATES = []string{
	"check_date_start",
	"check_date_end",
	"check_period_start",
	"check_period_end",
	//
	"docs_req_date",
	//
	"docs_memorandum_date",
	//
	"docs_act_date",
	//
	"docs_decision_date",
	//
	"docs_charge_date",
}

var CHECKS_KEY_INT = []string{
	"id",
	"inspector_id",
}

var ACTLIST_KEYS = map[string]string{
	"pay_year":          "int",
	"pay_month":         "int",
	"overpay_sicklist":  "float",
	"overpay_birth":     "float",
	"overpay_1_early":   "float",
	"underpay_sicklist": "float",
	"underpay_birth":    "float",
	"underpay_1_early":  "float",
	"ndfl_sicklist":     "float",
	"postal_expenses":   "float",
}

//JSON конвертирует time.Time в формат ISO, а нам нужна просто дата
func convertDates(data map[string]interface{}) {
	for _, k := range CHECKS_KEY_DATES {
		_, ok := data[k]
		if !ok {
			continue
		}
		v, ok := data[k].(time.Time)
		if !ok {
			log.Printf("Поле %s входит в список дат, но в данных не является типом time.Time: %v", k, data)
			continue
		}
		data[k] = v.Format(COMMON_DATE_LAYOUT)
	}
}

type TFilter struct {
	date1     string
	date2     string
	inspector int64
	regnum    string
}

func (dk *DBKeeper) SelectChecks(filter TFilter) DBResult {
	conditions := make([]string, 0)
	values := make(map[string]interface{})
	if filter.date1 > "" {
		conditions = append(conditions, "(check_date_start >= :date1 or check_date_end >= :date1)")
		values["date1"] = filter.date1
	}
	if filter.date2 > "" {
		conditions = append(conditions, "(check_date_start <= :date2 or check_date_end <= :date2)")
		values["date2"] = filter.date2
	}
	if filter.inspector >= 0 {
		conditions = append(conditions, "(inspector_id = :inspector)")
		values["inspector"] = filter.inspector
	}
	if filter.regnum > "" {
		conditions = append(conditions, "(insurer_regnum = :regnum)")
		values["regnum"] = filter.regnum
	}
	where := ""
	if len(conditions) > 0 {
		where = " where "
	}
	sql := "select " + strings.Join(CHECKS_KEY_LIST, ",") + " from checks " + where +
		strings.Join(conditions, " and ") +
		" order by check_date_start desc, id desc"
	var res []map[string]interface{} = make([]map[string]interface{}, 0)
	rows, err := dk.db.NamedQuery(sql, values)
	if err != nil {
		log.Printf("DBKeeper.SelectChecks: select error: %v\n", err)
		return DBResult{err, nil}
	}
	for rows.Next() {
		data := make(map[string]interface{})
		err = rows.MapScan(data)
		if err != nil {
			log.Printf("DBKeeper.SelectChecks: rowscan error: %v\n", err)
		} else {
			convertDates(data)
			res = append(res, data)
		}
	}
	rows.Close()
	return DBResult{nil, res}
}

func (dk *DBKeeper) GetCheckById(id int64, convertDate bool) DBResult {
	sql := "select " + strings.Join(CHECKS_KEY_LIST, ",") + " from checks where " +
		" id=$1 "
	var res map[string]interface{} = make(map[string]interface{})
	rows, err := dk.db.Queryx(sql, id)
	if err != nil {
		log.Printf("DBKeeper.GetCheckById: select error: %v\n", err)
		return DBResult{err, nil}
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.MapScan(res)
		if err != nil {
			log.Printf("DBKeeper.GetCheckById: rowscan error: %v\n", err)
			return DBResult{err, nil}
		}
	}
	if convertDate {
		convertDates(res)
	}
	return DBResult{nil, res}

}

func validateInput(data map[string]interface{}) error {
	//обязательно дожен быть ключ id
	_, ok := data["id"]
	if !ok {
		log.Printf("validateInput: нет id в данных %v\n", data)
		return fmt.Errorf("Нет id в данных")
	}
	//собираем неправильные ключи и сообщаем об этом
	invalid_keys := make([]string, 0)
	for k, _ := range data {
		valid := false
		for _, k1 := range CHECKS_KEY_LIST {
			if k == k1 {
				valid = true
			}
		}
		if !valid {
			log.Printf("validateInput: обнаружено неправильное поле %s\n", k)
			invalid_keys = append(invalid_keys, k)
		}
	}
	//удаляем неправильные ключи
	for _, k := range invalid_keys {
		delete(data, k)
	}

	//проверяем поля дат, неправильные отбрасываем
	for _, k := range CHECKS_KEY_DATES {
		_, ok := data[k]
		if !ok {
			continue
		}
		v, ok := data[k].(string) //на входе должна быть строка
		if !ok {
			delete(data, k)
			continue
		}
		_, err := time.Parse(COMMON_DATE_LAYOUT, v)
		if err != nil {
			delete(data, k)
		}
	}

	//стандартный парсер json->map[string]interface{}
	//все числа распознает как float64
	//конвертируем по списку в целое
	for _, k := range CHECKS_KEY_INT {
		var v int64
		_, ok := data[k]
		if !ok {
			continue
		}
		_v, ok := data[k].(float64)
		if ok {
			v = int64(_v)
		} else {
			v = 0
		}
		data[k] = v
	}
	return nil
}

//вставка/обновление проверки
func (dk *DBKeeper) UpsertCheck(data map[string]interface{}, command int) DBResult {
	err := validateInput(data)
	if err != nil {
		return DBResult{err, nil}
	}
	id := data["id"].(int64)
	//если id = 0, тогда вставка, иначе обновление
	if id == 0 {
		//insert
		flds1 := make([]string, 0)
		flds2 := make([]string, 0)
		for k, _ := range data {
			if k != "id" {
				flds1 = append(flds1, k)
				flds2 = append(flds2, ":"+k)
			}
		}
		sql := "insert into checks(" + strings.Join(flds1, ",") + ") values (" +
			strings.Join(flds2, ",") + ") returning id;"
		tx, err := dk.db.Beginx()
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Begin insert error: %v\n", err)
			return DBResult{fmt.Errorf("Ошибка начала транзакции"), nil}
		}
		stmt, err := tx.PrepareNamed(sql)
		err = stmt.Get(&id, data) //<-- в переменную id приходит идентификатор добавленной проверки
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Get insert error: %v\n", err)
		}
		err = tx.Commit()
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Commit insert error: %v\n", err)
		}
	} else {
		//update
		flds1 := make([]string, 0)
		for k, _ := range data {
			if k != "id" {
				flds1 = append(flds1, k+"=:"+k)
			}
		}
		sql := "update checks set " + strings.Join(flds1, ",") + " where id=:id;"
		tx, err := dk.db.Beginx()
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Begin update error: %v\n", err)
			return DBResult{fmt.Errorf("Ошибка начала транзакции"), nil}
		}
		_, err = tx.NamedExec(sql, data)
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Get update error: %v\n", err)
		}
		err = tx.Commit()
		if err != nil {
			log.Printf("DBKeeper.UpsertCheck: Commit update error: %v\n", err)
		}
	}
	if command == CMD_DOCS_REQ {
		dk.SetNumber(id, "docs_req_number", "docs_req_date")
	}
	if command == CMD_DOCS_MEMORANDUM {
		dk.SetNumber(id, "docs_memorandum_number", "docs_memorandum_date")
	}
	if command == CMD_DOCS_ACT {
		dk.SetNumber(id, "docs_act_number", "docs_act_date")
	}
	if command == CMD_DOCS_DECISION {
		dk.SetNumber(id, "docs_decision_number", "docs_decision_date")
	}
	if command == CMD_DOCS_CHARGE {
		dk.SetNumber(id, "docs_charge_number", "docs_charge_date")
	}

	return dk.GetCheckById(id, true)
}

func (dk *DBKeeper) SetNumber(id int64, field_number, field_date string) {
	sql := fmt.Sprintf(`with src as (select extract(year from %s) y from checks where id=$1),
	t as (select max(%s)+1 maxn from checks where extract(year from %s) in (select y from src))
	update checks set %s=t.maxn from t where id=$1 and %s=0;`, field_date, field_number, field_date,
		field_number, field_number)
	tx, err := dk.db.Beginx()
	if err != nil {
		log.Printf("DBKeeper.SetNumber(%d,%s,%s): Begin update error: %v\n", id, field_number, field_date, err)
		return
	}
	_, err = tx.Exec(sql, id)
	if err != nil {
		log.Printf("DBKeeper.SetNumber(%d,%s,%s): Get update error: %v\n", id, field_number, field_date, err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("DBKeeper.SetNumber(%d,%s,%s): Commit update error: %v\n", id, field_number, field_date, err)
	}

}

func json2int64(inp map[string]interface{}, key string) int64 {
	_, ok := inp[key]
	if !ok {
		log.Printf("json2int64: no key `%s' in data\n", key)
		return 0
	}
	tmp, ok := inp[key].(float64)
	if !ok {
		log.Printf("json2int64: cannot convert [%s]=`%#v' to int\n", key, inp[key])
		return 0
	}
	/*
		res, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			log.Printf("json2int64: cannot convert [%s]=`%s' to int\n", key, tmp)
			return 0
		}*/
	return int64(tmp)
}

func json2float64(inp map[string]interface{}, key string) float64 {
	_, ok := inp[key]
	if !ok {
		log.Printf("json2float64: no key `%s' in data\n", key)
		return 0
	}
	res := inp[key].(float64)
	if !ok {
		log.Printf("json2float64: cannot convert [%s]=`%#v' to float64\n", key, inp[key])
		return 0.0
	}

	/*
		res, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			log.Printf("json2float64: cannot convert [%s]=`%s' to float\n", key, tmp)
			return 0
		}*/
	return res
}

//вставка списка для акта
func (dk *DBKeeper) UpsertActList(data map[string]interface{}) DBResult {
	id := json2int64(data, "id")
	if id == 0 {
		return DBResult{fmt.Errorf("id=0"), nil}
	}
	//prepare data
	oldrows := data["rows"].([]interface{})
	rows := make([]map[string]interface{}, 0)
	for _, obj := range oldrows {
		cnvobj := obj.(map[string]interface{})
		newobj := make(map[string]interface{})
		newobj["id_check"] = id
		doadd := false
		for key, typ := range ACTLIST_KEYS {
			if typ == "int" {
				newobj[key] = json2int64(cnvobj, key)
			}
			if typ == "float" {
				v := json2float64(cnvobj, key)
				if v > 0 {
					doadd = true
				}
				newobj[key] = v
			}
		}
		if doadd {
			rows = append(rows, newobj)
		}
	}

	tx, err := dk.db.Beginx()
	if err != nil {
		log.Printf("DBKeeper.UpsertActList: Begin tx error: %v\n", err)
		return DBResult{fmt.Errorf("Ошибка начала транзакции"), nil}
	}
	_, err = tx.Exec("delete from act_list where id_check=$1", id)
	if err != nil {
		tx.Rollback()
		log.Printf("DBKeeper.UpsertActList: delete error: %v\n", err)
		return DBResult{fmt.Errorf("Ошибка удаления"), nil}
	}
	for _, row := range rows {
		flds1 := make([]string, 0)
		flds2 := make([]string, 0)
		for k, _ := range row {
			flds1 = append(flds1, k)
			flds2 = append(flds2, ":"+k)
		}
		sql := "insert into act_list(" + strings.Join(flds1, ",") + ") values (" +
			strings.Join(flds2, ",") + ") on conflict do nothing;"
		stmt, err := tx.PrepareNamed(sql)
		if err != nil {
			tx.Rollback()
			log.Printf("DBKeeper.UpsertActList: prepare insert error: %v\n", err)
			return DBResult{fmt.Errorf("Ошибка вставки"), nil}
		}
		_, err = stmt.Exec(row)
		if err != nil {
			tx.Rollback()
			log.Printf("DBKeeper.UpsertActList: insert error: %v\n", err)
			return DBResult{fmt.Errorf("Ошибка вставки"), nil}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("DBKeeper.UpsertActList: commit error: %v\n", err)
		return DBResult{fmt.Errorf("Ошибка вставки"), nil}
	}

	return dk.GetActList(id)
}

func (dk *DBKeeper) GetActList(id int64) DBResult {
	sql := `select * from act_list where id_check=$1 order by pay_year, pay_month`
	var res []map[string]interface{} = make([]map[string]interface{}, 0)
	rows, err := dk.db.Queryx(sql, id)
	if err != nil {
		log.Printf("DBKeeper.GetActList: select error: %v\n", err)
		return DBResult{err, nil}
	}
	for rows.Next() {
		data := make(map[string]interface{})
		err = rows.MapScan(data)
		if err != nil {
			log.Printf("DBKeeper.GetActList: rowscan error: %v\n", err)
		} else {
			res = append(res, data)
		}
	}
	rows.Close()
	return DBResult{nil, res}

}
