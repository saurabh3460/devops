package main

import (
	"fmt"
	"sync"
)

/*
Th eidea is that the program will give a result based on the performances of the players of each
team (a json file). Each player (player.go file) will write in a channel the amout of
yellow / red cards he got from the game, based on his agressvity and a random number generated.
And the game.go file will read from the channels every time and update the result.
*/

var teamA = team{name: "a", players: []players{
	{number: 1}, {number: 2}, {number: 3}, {number: 4},
}}

type players struct {
	number int
	yellow int
	red    int
}

type team struct {
	name    string
	players []players
}

type Performance struct {
	// based on agressvity and a generated number
	yellow int
	red    int
}

// a player func to write to performance channel
// a game func to read from channel to update result and update players yellow and red

func player(performance chan Performance, wg *sync.WaitGroup) {
	// wg.Add(1)
	fmt.Println("player operation....")
	performance <- Performance{yellow: 1, red: 2}
	wg.Done()
}

func game(performance chan Performance, wg *sync.WaitGroup) {
	// wg.Add(1)
	fmt.Println("game operation....")
	input := <-performance
	fmt.Println(input)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	performance := make(chan Performance)
	wg.Add(1)
	player(performance, &wg)
	wg.Add(1)
	game(performance, &wg)
	wg.Wait()
	fmt.Println("all done ")
}
