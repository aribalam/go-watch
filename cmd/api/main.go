package main

import (
	"backend/models"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version string = "1.0.1"

type Config struct {
	Port int
	Env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Application struct {
	config Config
	logger *log.Logger
	models models.Models
}

func main() {

	var config Config
	flag.IntVar(&config.Port, "port", 4000, "Port Number")
	flag.StringVar(&config.Env, "env", "development", "Environment")
	flag.StringVar(&config.db.dsn, "dsn", "postgres://postgres:180eedcd@localhost:5432/go_movies?sslmode=disable", "postgres connection")
	flag.Parse()

	db, err := openDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	application := &Application{
		config: config,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", application.config.Port),
		Handler:      application.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.logger.Println("Listening on Port", application.config.Port)
	err = srv.ListenAndServe()
	if err != nil {
		application.logger.Println(err)
	}
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
