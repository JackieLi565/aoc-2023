package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MaxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Cube1() {

		data, err := DataParser()

		if (err != nil) {
			return
		}
		total := 0

		Reset:
			for i, game := range data {
				for _, bag := range game {
					if bag.amount > MaxCubes[bag.color] {
						continue Reset
					}
				}
				total += i + 1
			}

		fmt.Println(total)
}

type Data struct{
	color  string
	amount int
}

func DataParser()([][]Data, error) {
	file, err := os.Open("./data.txt");
	data := [][]Data{}

	if err != nil {
		return data, errors.New("failed to open file") 
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")[1]
		currentGame := strings.ReplaceAll(line, "; ", ", ")
		bags := strings.Split(currentGame, ", ")
		temp := []Data{}

		for _, bag := range bags {
			values := strings.Split(bag, " ")
			amount, err := strconv.Atoi(values[0])
			color := values[1]

			if err != nil {
				return data, errors.New("failed to convert string to number")
			}
			temp = append(temp, Data{color: color, amount: amount})
		}
		data = append(data, temp)
	}

	return data, nil
}
