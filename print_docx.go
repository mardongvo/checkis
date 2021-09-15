package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const PRINT_DATE_LAYOUT = "02.01.2006"

var ALIASES = map[string]string{
	"insurer_regnum":      "Регистрационный_номер",
	"insurer_kpsnum":      "Код_подчиненности",
	"insurer_inn":         "ИНН",
	"insurer_kpp":         "КПП",
	"insurer_fullname":    "Полное_наименование_страхователя",
	"insurer_shortname":   "Сокращенное_наименование_страхователя",
	"insurer_postindex":   "Индекс_страхователя",
	"insurer_postaddress": "Адрес_страхователя",
}

var MONTHS = map[int64]string{
	0:  "-",
	1:  "Январь",
	2:  "Февраль",
	3:  "Март",
	4:  "Апрель",
	5:  "Май",
	6:  "Июнь",
	7:  "Июль",
	8:  "Август",
	9:  "Сентябрь",
	10: "Октябрь",
	11: "Ноябрь",
	12: "Декабрь",
}

type ReplacementInfo struct {
	commonFields []string
	markers      map[string][][]string
	//marker       string
	//listFields   [][]string
}

func NewReplacementInfo() ReplacementInfo {
	var res ReplacementInfo
	res.commonFields = make([]string, 0)
	res.markers = make(map[string][][]string)
	//res.listFields = make([][]string, 0)
	return res
}

func (ri *ReplacementInfo) AddCommonField(key string, value string) {
	ri.commonFields = append(ri.commonFields, "#"+key+"#", value)
	alias, ok := ALIASES[key]
	if ok {
		ri.commonFields = append(ri.commonFields, "#"+alias+"#", value)
	}
}

func (ri *ReplacementInfo) AddMarker(marker string) {
	ri.markers[marker] = make([][]string, 0)
}

func (ri *ReplacementInfo) AddListReplacement(marker string, list []string) {
	ri.markers[marker] = append(ri.markers[marker], list)
}

func ApiPrintDoc(w http.ResponseWriter, r *http.Request) {
	var check_id, inspector_id, signer_id int64
	var cook *http.Cookie
	var err error
	var docx DocxTemplate
	var filename string

	vars := mux.Vars(r)
	//check_id
	check_id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		check_id = 0
	}
	//template_id
	template_id := vars["template"]
	//inspector_id
	cook, err = r.Cookie("inspector")
	if err != nil {
		inspector_id = 0
	} else {
		inspector_id, err = strconv.ParseInt(cook.Value, 10, 64)
		if err != nil {
			inspector_id = 0
		}
	}
	//signer_id
	cook, err = r.Cookie("signer")
	if err != nil {
		signer_id = 0
	} else {
		signer_id, err = strconv.ParseInt(cook.Value, 10, 64)
		if err != nil {
			signer_id = 0
		}
	}
	//
	replacementInfo := dbkeeper.PreparePrintInfo(check_id, inspector_id, signer_id)
	if template_id == "req" {
		docx, err = OpenTemplate("./template/template_docs_req.docx")
		filename = "trebovanie.docx"
	}
	if template_id == "memorandum" {
		docx, err = OpenTemplate("./template/template_docs_memorandum.docx")
		filename = "doklzapiska.docx"
	}
	if template_id == "act" {
		docx, err = OpenTemplate("./template/template_docs_act.docx")
		filename = "akt.docx"
	}
	if template_id == "decision" {
		docx, err = OpenTemplate("./template/template_docs_decision.docx")
		filename = "reshenie.docx"
	}
	if template_id == "charge" {
		docx, err = OpenTemplate("./template/template_docs_charge.docx")
		filename = "trebovaine_uplata.docx"
	}
	if err != nil {
		log.Printf("ApiPrintDoc: OpenTemplate %v\n", err)
		return
	}
	defer docx.Close()
	docx.Replace(replacementInfo.commonFields)
	docx.ReplaceToList("p", "#doc_req#", replacementInfo.markers["#doc_req#"])
	docx.ReplaceToList("tr", "#actlist_month#", replacementInfo.markers["#actlist_month#"])
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	err = docx.Write(w)
	if err != nil {
		log.Printf("ApiPrintDoc: write error %v", err)
		return
	}
}

func (dk *DBKeeper) PreparePrintInfo(check_id int64, inspector_id int64, signer_id int64) ReplacementInfo {
	var err error
	var res ReplacementInfo = NewReplacementInfo()
	dbres := dk.GetCheckById(check_id, false)
	if dbres.Error != nil {
		return res
	}
	for k, v := range dbres.Data.(map[string]interface{}) {
		switch vv := v.(type) {
		case string:
			res.AddCommonField(k, vv)
		case int64:
			res.AddCommonField(k, strconv.FormatInt(vv, 10))
		case float64:
			res.AddCommonField(k, strconv.FormatFloat(vv, 'f', 2, 64))
		case time.Time:
			res.AddCommonField(k, vv.Format(PRINT_DATE_LAYOUT))
		}
	}
	//
	inspector_info := StaffInfo{}
	err = dk.db.Get(&inspector_info, "select * from staff where id=$1", inspector_id)
	if err != nil {
		log.Printf("PreparePrintInfo: %v\n", err)
	}
	res.AddCommonField("Исполнитель_ФИО", inspector_info.Fio)
	res.AddCommonField("Исполнитель_Должность", inspector_info.Position)
	res.AddCommonField("Исполнитель_Телефон", inspector_info.Phone)

	signer_info := StaffInfo{}
	err = dk.db.Get(&signer_info, "select * from staff where id=$1", signer_id)
	if err != nil {
		log.Printf("PreparePrintInfo: %v\n", err)
	}
	res.AddCommonField("Руководитель_ФИО", signer_info.Fio)
	res.AddCommonField("Руководитель_Должность", signer_info.Position)
	res.AddCommonField("Руководитель_ФИО_Дат", signer_info.Fio_Dative)
	res.AddCommonField("Руководитель_Должность_Дат", signer_info.Position_Dative)
	res.AddCommonField("Руководитель_ФИО_Тв", signer_info.Fio_Genitive)
	res.AddCommonField("Руководитель_Должность_Тв", signer_info.Position_Genitive)

	//doc_req replacements
	sqlprep := make([]string, 0)
	for _, fld := range []string{
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
	} {
		sqlprep = append(sqlprep, fmt.Sprintf(`select id, doc_title from docs_list where flag='%s'
	and exists (select * from checks where %s=1 and id=$1) `, fld, fld))
	}
	sql := "select doc_title from (" + strings.Join(sqlprep, "union\n") + ") a order by id;"
	doc_titles := make([]string, 0)
	err = dk.db.Select(&doc_titles, sql, check_id)
	if err != nil {
		log.Printf("PreparePrintInfo: %v\n", err)
	}
	res.AddMarker("#doc_req#")
	for i, dc := range doc_titles {
		list := []string{"#n#", strconv.FormatInt(int64(i+1), 10), "#doc_req#", dc}
		res.AddListReplacement("#doc_req#", list)
	}

	//act_list replacements
	overpay_all := float64(0)
	underpay_all := float64(0)
	ndfl_all := float64(0)
	postal_all := float64(0)
	for i, fld := range []string{
		"overpay_sicklist",
		"overpay_birth",
		"overpay_1_early",
		"underpay_sicklist",
		"underpay_birth",
		"underpay_1_early",
		"ndfl_sicklist",
		"postal_expenses",
	} {
		var tmp float64
		err = dk.db.Get(&tmp, fmt.Sprintf("select sum(coalesce(%s,0.0)) from act_list where id_check=$1;", fld), check_id)
		if err != nil {
			log.Printf("PreparePrintInfo: act_list[%s] = %v\n", fld, err)
		}
		switch i {
		case 0, 1, 2:
			overpay_all += tmp
		case 3, 4, 5:
			underpay_all += tmp
		case 6:
			ndfl_all += tmp
		case 7:
			postal_all += tmp
		}
		res.AddCommonField(fld, strings.ReplaceAll(fmt.Sprintf("%.2f", tmp), ".", ","))
	}
	res.AddCommonField("overpay_all", strings.ReplaceAll(fmt.Sprintf("%.2f", overpay_all), ".", ","))
	res.AddCommonField("underpay_all", strings.ReplaceAll(fmt.Sprintf("%.2f", underpay_all), ".", ","))
	res.AddCommonField("ndfl_all", strings.ReplaceAll(fmt.Sprintf("%.2f", ndfl_all), ".", ","))
	res.AddCommonField("postal_all", strings.ReplaceAll(fmt.Sprintf("%.2f", postal_all), ".", ","))
	list_by_month := []struct {
		Pay_year  int64
		Pay_month int64
		ASum      float64
		Ndfl      float64
		Postal    float64
	}{}
	sql = `select coalesce(pay_year,0) pay_year, coalesce(pay_month,0) pay_month,
	 coalesce(overpay_sicklist,0.0)+coalesce(overpay_birth,0.0)+coalesce(overpay_1_early,0.0) asum,
	 coalesce(ndfl_sicklist, 0.0) ndfl,
	 coalesce(postal_expenses, 0.0) postal
	 from act_list where id_check=$1 order by pay_year, pay_month;`
	err = dk.db.Select(&list_by_month, sql, check_id)
	if err != nil {
		log.Printf("PreparePrintInfo: %v\n", err)
	}
	res.AddMarker("#actlist_month#")
	for _, e := range list_by_month {
		list := []string{
			"#actlist_year#", fmt.Sprintf("%d", e.Pay_year),
			"#actlist_month#", MONTHS[e.Pay_month],
			"#actlist_sum#", strings.ReplaceAll(fmt.Sprintf("%.2f", e.ASum), ".", ","),
			"#actlist_ndfl#", strings.ReplaceAll(fmt.Sprintf("%.2f", e.Ndfl), ".", ","),
			"#actlist_postal#", strings.ReplaceAll(fmt.Sprintf("%.2f", e.Postal), ".", ","),
		}
		res.AddListReplacement("#actlist_month#", list)
	}
	return res
}
