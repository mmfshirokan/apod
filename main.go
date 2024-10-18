package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mmfshirokan/apod/docs"
	"github.com/mmfshirokan/apod/internal/config"
	"github.com/mmfshirokan/apod/internal/consumer"
	"github.com/mmfshirokan/apod/internal/handlers"
	"github.com/mmfshirokan/apod/internal/repository"
	"github.com/mmfshirokan/apod/internal/service"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title APOD Serevr
// @version 1.0
// @description This is a server for storing NASA apod data & images.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cnf, err := config.New()
	if err != nil {
		log.Fatal("Fatal parse config error: ", err)
	}

	pgDB, err := pgxpool.New(ctx, cnf.PostgresURL)
	if err != nil {
		log.Fatal("Fatal MySQl connection error: ", err)
	}
	defer pgDB.Close()

	infoRepo := repository.NewInfo(pgDB)
	infoServ := service.NewInfo(infoRepo)

	imgRepo := repository.NewImage(cnf.ImageDestenation)
	imageServ := service.NewImage(imgRepo)

	h := handlers.New(infoServ, cnf.NginxURL)
	c := consumer.New(infoServ, imageServ)

	go c.Consume(ctx, cnf.Target, cnf.ApiKey)

	router := httprouter.New()

	router.GET("/get/:date", h.Get)
	router.GET("/get", h.GetAll)
	router.GET("/docs/:any", swaggerHandler)

	server := &http.Server{
		Addr:    ":" + cnf.ServerPort,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Fatal server error: ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	go func() {
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ch

		cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Fatal server shutdown error: ", err)
		}
	}()

	wg.Wait()
	log.Info("Server gracefully stopped")
}

func swaggerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(w, r)
}
