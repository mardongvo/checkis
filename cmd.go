package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var dbkeeper DBKeeper

func UIHandle(filepath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(filepath, "./template/common_header.tmpl")
		if err != nil {
			log.Printf("Ошибка парсинга шаблона (%s): %v", filepath, err)
			fmt.Printf("Ошибка парсинга шаблона (%s): %v\n", filepath, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Template parsing error"))
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Ошибка рендеринга шаблона (%s): %v", filepath, err)
			fmt.Printf("Ошибка рендеринга шаблона (%s): %v", filepath, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Template rendering error"))
			return
		}
	}

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered: %v", r)
		}
	}()
	//config
	cfgfile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Ошибка при открытии файла конфигурации: %v", err)
	}
	cfg, err := ReadConfig(cfgfile)
	cfgfile.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//log
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if cfg.LogFile > "" {
		logfile, err := os.Create(cfg.LogFile)
		if err != nil {
			log.Fatalf("Невозможно создать файл: %v", err)
		}
		defer logfile.Close()
		log.SetOutput(logfile)
	}

	dbkeeper = NewDBKeeper(cfg.Database)

	mixer := mux.NewRouter()
	mixer.HandleFunc("/", UIHandle("./template/root.tmpl"))
	mixer.HandleFunc("/check", UIHandle("./template/form_check.tmpl"))
	mixer.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))
	//
	mixer.HandleFunc("/api/insurers/{regnum:[0-9]+}", ApiFindInsurer).Methods("GET")
	mixer.HandleFunc("/api/staff", ApiStaffList).Methods("GET")
	//
	mixer.HandleFunc("/api/checks", ApiListChecks).Methods("GET")
	mixer.HandleFunc("/api/checks/{id:[0-9]+}", ApiGetCheck).Methods("GET")
	mixer.HandleFunc("/api/checks/{id:[0-9]+}/{part:(?:actlist)}", ApiGetPart).Methods("GET")
	mixer.HandleFunc("/api/checks/{id:[0-9]+}/{part:(?:header|req|memorandum|act|actlist|decision|charge)}", ApiSaveCheck).Methods("POST")
	mixer.HandleFunc("/print/{template:(?:req|memorandum|act|decision|charge)}/{id:[0-9]+}", ApiPrintDoc).Methods("GET")

	srv := &http.Server{
		Addr:           cfg.ListenAddress,
		Handler:        mixer,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
