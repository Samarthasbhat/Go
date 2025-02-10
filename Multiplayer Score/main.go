// Requirement
//Players scoreboard
// ScoreManager
// // Features
// Add new players to the scoreboard.
//  Update player scores concurrently.
// Retrieve and display the scoreboard in real-time.
// Handle errors like duplicate player entries or invalid score updates.
// Use concurrency (goroutines and channels) to handle multiple score updates efficiently.

package main

import (
	"errors"
	"fmt"
)

type ScoreManager interface{
	AddPlayer(playerID int, PlayerName string) error
	UpdateScore(playerID int, score float64) error
	fetchScore(playerID int) (float64, error)
}

type Player struct{
	PlayerName string
	PlayerID int
}

type ScoreBoard struct{
	players map[int]string
	score map[int]float64

}


func (s *ScoreBoard) AddPlayer(playerID int, playerName string) error{
	if _, exists := s.players[playerID]; exists{
		return errors.New("Player already Exists")
	}
	s.players[playerID] = playerName
	s.score[playerID] = 0
	return nil
}


func main(){
	scoreboard := ScoreBoard{
		players: make(map[int]string),
		score: make(map[int]float64),
	}

	add := scoreboard.AddPlayer(1, "Ram")
	if add != nil{
		fmt.Println("Error:", add)
	}else{
		fmt.Println("Player Added successfully")
	}

	// add = scoreboard.AddPlayer(1,"shayn")
	// if add != nil {
	// 	fmt.Println("Error:", add)
	// }
}
