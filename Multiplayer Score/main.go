// Requirement
//Players scoreboard
// ScoreManager
// // Features
// Add new players to the scoreboard.
//  Update player scores concurrently.
// Retrieve and display the scoreboard in real-time.
// Handle errors like duplicate player entries or invalid score updates.
// Use concurrency (goroutines and channels) to handle multiple score updates efficiently.
// 1️⃣ Uses a channel (scoreUpdateChan) for non-blocking, concurrent score updates.
// 2️⃣ A dedicated goroutine (processScoreUpdates()) listens for and applies updates efficiently.
// 3️⃣ Eliminates race conditions and improves scalability by reducing direct locking.
// 4️⃣ Supports real-time scoreboard display with a graceful shutdown mechanism.


package main

import (
	"errors"
	"fmt"
	"sync"
)

type ScoreManager interface{
	AddPlayer(playerID int, PlayerName string) error
	UpdateScore(playerID int, score float64) error
	FetchScore(playerID int) (float64, error)
	DisplayScoreboard()
}

type Player struct{
	PlayerName string
	PlayerID int
}

type ScoreBoard struct{
	players map[int]string
	score map[int]float64
	scoreUpdateChan  chan  scoreUpdateRequest
	mu sync.RWMutex
}

// Struct for score update request
type scoreUpdateRequest struct{
	playerID int
	score float64
}

// NewScoreBoard initializes a new scoreboard
func NewScoreBoard() *ScoreBoard{
	sb := &ScoreBoard{
		players: make(map[int]string),
		score: make(map[int]float64),
		scoreUpdateChan: make(chan scoreUpdateRequest, 10),
	}
	// Start processing updates in a goroutine
	go sb.processScoreUpdates()
	return sb
}

// Goroutine to process score updates
func (s *ScoreBoard) processScoreUpdates(){
	for update := range s.scoreUpdateChan{
		s.mu.Lock()
		if _, exists := s.players[update.playerID]; exists{
			s.score[update.playerID] += update.score
		}else{
			fmt.Printf("Error: Player %d not found\n", update.playerID)
		}
		s.mu.Unlock()
	}
}

// AddPlayer adds a new player to the scoreboard
func (s *ScoreBoard) AddPlayer(playerID int, playerName string) error{
	
	s.mu.Lock()
	defer s.mu.Unlock()

	
	if _, exists := s.players[playerID]; exists{
		return errors.New("Player already Exists")
	}
	s.players[playerID] = playerName
	s.score[playerID] = 0
	return nil
}

// UpdateScore updates a player's score concurrently
func (s *ScoreBoard) UpdateScore(playerID int, score float64) error{
	s.mu.RLock()

	if _,exists := s.players[playerID]; !exists {
		return errors.New("Players not found")
	}
	s.mu.RUnlock()

	// Send update request to channel
	s.scoreUpdateChan <- scoreUpdateRequest{playerID, score}
	return nil
}

// FetchScore retrieves a player's score
func (s *ScoreBoard) FetchScore(playerID int) (float64,error){
	s.mu.RLock()
	defer s.mu.RUnlock()

	score, exists := s.score[playerID]
	if !exists {
		return 0, errors.New("Player not found")
	}
	return score, nil
}

// DisplayScoreboard prints the scoreboard
func (s *ScoreBoard) DisplayScoreboard(){
	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Println("==== Scoreboard ====")
	for id, name := range s.players{
		fmt.Printf("%s (%d): %.2f\n", name, id, s.score[id])
	}
	fmt.Println("================")
}

// Close the score update channel to stop processing
func (s *ScoreBoard) Stop(){
	close(s.scoreUpdateChan)
}

func main(){
	scoreboard := NewScoreBoard()

	// Adding players
	
	if add := scoreboard.AddPlayer(1, "Ram"); add != nil{
		fmt.Println("Error:", add)
	}else{
		fmt.Println("Player Added successfully")
	}

	if add := scoreboard.AddPlayer(2,"shan"); add != nil{
		fmt.Println("Error:", add)
	}else{
		fmt.Println("player shan added successfully")
	}


	// add = scoreboard.AddPlayer(1,"shayn")
	// if add != nil {
	// 	fmt.Println("Error:", add)
	// }

	// Concurrent score updates using goroutines

	var wg sync.WaitGroup
	scoreUpdates := []struct{
		playerID int
		score float64
	}{
		{1, 10.5}, {2, 5.0}, {1, 15.0}, {2, 7.5}, {1, 20.0},
	}

	for _, update := range scoreUpdates{
		wg.Add(1)
		go func(playerID int, score float64){
		defer wg.Done()
		if add := scoreboard.UpdateScore(playerID, score); add != nil{
			fmt.Println("Error:",add)
		}
	}(update.playerID, update.score)
}

wg.Wait()

// Display scoreboard
scoreboard.DisplayScoreboard()

// Stop processing
scoreboard.Stop()
}
