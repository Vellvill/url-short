package main

import (
	"NewOne/config"
	service "NewOne/internal"
	"NewOne/internal/postgres"
	"NewOne/internal/repository"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	cfg := config.GetConfig()

	client, err := postgres.NewClient(context.TODO(), cfg.Storage)
	if err != nil {
		log.Fatal()
	}

	repo := repository.NewRepository(client)

	impl := service.New(repo)

	startSt := make(chan struct{})
	go startStatus(startSt)
	go repo.Status(context.TODO(), startSt)

	http.HandleFunc("/add", impl.AddNewUrl)

	http.HandleFunc("/", impl.RedirectToUrl)

	http.HandleFunc("/getstats/", impl.CheckStats)

	http.HandleFunc("/checkstatus", impl.CheckStatus)

	start(cfg)
}

func start(cfg *config.Config) {
	listner, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	if err != nil {
		log.Fatal()
	}
	if err = http.Serve(listner, nil); err != nil {
		log.Fatal()
	}
	fmt.Printf("server is listening %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
}

func startStatus(startSt chan struct{}) {
	startSt <- struct{}{}
	for {
		select {
		case <-time.After(10 * time.Minute):
			startSt <- struct{}{}
		}
	}
}
