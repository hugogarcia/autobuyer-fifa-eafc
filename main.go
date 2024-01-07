package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/auth"
	"github.com/hugogarcia/fifa-auto-buyer/entity"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/search"
	"github.com/joho/godotenv"
)

func main() {
	err := logger.SetLogFile("./logs.log")
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	entity.LoadPlayersFromCSV()

	auth.SetNewToken()
	//db.ReadJsonFile()
	go search.SearchPlayers()
	time.Sleep(time.Second * 30)
	go search.GetWatchlistAndRebid()

	listen()
}

func listen() {
	sig := make(chan os.Signal, 500)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	signal.Notify(sig, os.Interrupt, syscall.SIGABRT)
	<-sig
	logger.LogMessage(nil, time.Now().String()+" - CLOSED")
}
