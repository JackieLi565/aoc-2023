package main

import (
	"fmt"
)

func main() {
	data, err := DataParser()

	if err != nil {
		return
	}
	sumPower := 0

	for _, game := range data {
		gamePower := 1
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, bag := range game {
			if bag.amount > minCubes[bag.color] {
				minCubes[bag.color] = bag.amount
			}
		}

		for _, amount := range minCubes {
			gamePower *= amount
		}

		sumPower += gamePower
	}

	fmt.Println(sumPower)
}