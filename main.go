package main

import (
	"bareksa-take-home-test-michael-koh/config"
	newsService "bareksa-take-home-test-michael-koh/core/service/news"
	"bareksa-take-home-test-michael-koh/handler"
	repoNews "bareksa-take-home-test-michael-koh/repository/news"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// init config
	cfg := config.Get()

	// init db
	dsn := fmt.Sprintf(cfg.MySQLDSNFormat, cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Init DB error : %s\n", err.Error())
	}

	// init redis cache

	// init repo
	newsRepo := repoNews.NewRepository(db)

	// init service
	newsService := newsService.NewService(newsRepo)

	// init handler
	bareksaNewsHandler := handler.NewBareksaNewsHandler(newsService)

	// setup server and register route
	route := mux.NewRouter()
	route.HandleFunc("/", bareksaNewsHandler.GetNewsHandler)
	route.HandleFunc("/news", bareksaNewsHandler.GetNewsHandler).
		Queries("status", "{status}").
		Queries("topicSerials", "{topicSerials}").
		Methods("GET")

	http.Handle("/", route)

	srv := &http.Server{
		Handler:      route,
		Addr:         fmt.Sprintf("%s:%s", cfg.HostAdress, cfg.RESTPort),
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
	}

	log.Printf("Listening on Address : %s", cfg.HostAdress)
	log.Printf("Listening on port : %s", cfg.RESTPort)
	log.Printf("Write timeout : %d", cfg.WriteTimeout)
	log.Printf("Read timeout : %d", cfg.ReadTimeout)
	log.Fatal(srv.ListenAndServe())
}
