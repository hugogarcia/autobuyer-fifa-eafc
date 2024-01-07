package entity

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var TradePlayerMap = map[uint64]uint64{}

type Player struct {
	MaxBid uint64
	Name   string
}

var Players = map[uint64]Player{}

func LoadPlayersFromCSV() {
	f, err := os.Open("./players.csv")
	if err != nil {
		log.Panic(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	for i, v := range records {
		if i == 0 {
			continue
		}

		if len(v) < 3 {
			fmt.Printf("Problem on line %d, missing info \n", i)
			continue
		}

		playerId := v[0]
		playerName := v[1]
		playerBid := v[2]

		if playerId == "" || playerName == "" || playerBid == "" {
			fmt.Printf("Problem on line %d, missing info ID or MAX BID\n", i)
		}

		id, _ := strconv.ParseUint(playerId, 10, 64)
		bid, _ := strconv.ParseUint(playerBid, 10, 64)

		Players[id] = Player{
			Name:   playerName,
			MaxBid: bid,
		}
	}

	fmt.Println("PLAYERS LOADED")
}
