package main

import (
	"bareksa-take-home-test-michael-koh/config"
	serviceNews "bareksa-take-home-test-michael-koh/core/service/news"
	serviceTag "bareksa-take-home-test-michael-koh/core/service/tag"
	serviceTopic "bareksa-take-home-test-michael-koh/core/service/topic"
	"bareksa-take-home-test-michael-koh/handler"
	repoNews "bareksa-take-home-test-michael-koh/repository/news"
	repoTag "bareksa-take-home-test-michael-koh/repository/tag"
	repoTopic "bareksa-take-home-test-michael-koh/repository/topic"
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
	topicRepo := repoTopic.NewRepository(db)
	tagRepo := repoTag.NewRepository(db)

	// init service
	newsService := serviceNews.NewService(newsRepo)
	topicService := serviceTopic.NewService(topicRepo)
	tagService := serviceTag.NewService(tagRepo)

	// init handler
	bareksaNewsHandler := handler.NewBareksaNewsHandler(newsService, topicService, tagService)

	// setup server and register route
	route := mux.NewRouter()
	route.HandleFunc("/", bareksaNewsHandler.GetNewsHandler)

	// News Handler
	route.HandleFunc("/news", bareksaNewsHandler.GetNewsHandler).
		Queries("status", "{status}").
		Queries("topicSerials", "{topicSerials}").
		Methods("GET")
	route.HandleFunc("/news", bareksaNewsHandler.CreateNewsHandler).
		Methods("POST")
	route.HandleFunc("/news", bareksaNewsHandler.UpdateNewsHandler).
		Methods("PATCH")
	route.HandleFunc("/news/{newsSerial}", bareksaNewsHandler.DeleteNewsHandler).
		Methods("DELETE")

	// tags handler
	route.HandleFunc("/tags", bareksaNewsHandler.GetTagsHandler).
		Queries("page", "{page}").
		Queries("size", "{size}").
		Methods("GET")
	route.HandleFunc("/tags", bareksaNewsHandler.CreateTagsHandler).
		Methods("POST")
	route.HandleFunc("/tags/{tagName}", bareksaNewsHandler.UpdateTagsHandler).
		Methods("PUT")
	route.HandleFunc("/tags/{tagName}", bareksaNewsHandler.DeleteTagsHandler).
		Methods("DELETE")

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
