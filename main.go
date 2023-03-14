package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hugogarcia/fifa-auto-buyer/auth"
	"github.com/hugogarcia/fifa-auto-buyer/logger"
	"github.com/hugogarcia/fifa-auto-buyer/search"
)

func main() {
	err := logger.SetLogFile("./logs.log")
	if err != nil{
		panic(err);
	}

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
