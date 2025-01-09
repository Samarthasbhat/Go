package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type GameState struct{
		CurrentRoom string
		HasKey bool
}

func main(){
	game := GameState{
		CurrentRoom: "start",
		HasKey: 	false,
	}

	scanner := bufio.NewScanner(os.Stdin) // Use bufio.Scanner for multi-word input


	for {
		printRoom(game.CurrentRoom)

		fmt.Print("What do you want to do? ")
		scanner.Scan() // Read the entire line of input
		input := strings.ToLower(scanner.Text())


		//Handle commands
		switch input {
		case "look":
			lookAround(game.CurrentRoom)
		case "go north":
			if game.CurrentRoom == "start"{
				game.CurrentRoom = "room2"
			}else if game.CurrentRoom == "room2" && game.HasKey {
				fmt.Println("You unlocked the door with the key!")
				fmt.Println("Congratulations, you won!")
				return 
			}else{
				fmt.Println("The door is locked")
			}
		case "pick up key":
			if game.CurrentRoom == "room2" && !game.HasKey{
				game.HasKey = true
				fmt.Println("You picked up the key")
			}else if game.HasKey{
				fmt.Println("You already have the key.")
			}else{
				fmt.Println("There's is nothing to pick up here")
			}
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("I don't understand that command.")
		}
	}
}


func printRoom(room string){
	switch room {
	case "start":
		fmt.Println("You are in a small, dimly light room. There's a door to the north.")
	case "room2":
		fmt.Println("You are in another room. There's a locked door to the north, and you see a key on the floor")
		
	}
}


func lookAround(room string){
	switch room{
	case "start":
		fmt.Println("The room is small and plain. There's a door to the north")
	case "room2":
		fmt.Println("There is a key on the floor and a locked door to the north")
	}
}