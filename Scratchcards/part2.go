package main

import (
	"fmt"
	"log"
)

func Part2() {
	data, err := DataParser()

	if err != nil {
		log.Fatal(err)
		return
	}

	res := 0
	for i := 0; i < len(data); i++ {
		res += getRecursiveCards(i, data) + 1
	}
	
	fmt.Println(res)
}

func getNumberOfWinners(game Card) (wins int) {
	for _, num := range game.have {
		_, ok := game.win[num]
		if ok {
			wins++
		}
	}

	return
}

// bfs sum of wins per card instance
func getRecursiveCards(lookup int, cards []Card) int {
	wins := getNumberOfWinners(cards[lookup])
	acc := 0
	for i := 0; i < wins; i++ {
		acc += getRecursiveCards(i+lookup+1, cards)
	}

	return acc + wins
}