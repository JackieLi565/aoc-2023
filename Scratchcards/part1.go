package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	data, err := DataParser()

	if err != nil {
		return
	}
	totalSum := 0
	for _, game := range data {
		temp := 0
		hasFoundOne := false
		for _, num := range game.have {
			_, ok := game.win[num]
			if ok {
				if !hasFoundOne {
					temp = 1
					hasFoundOne = true
				} else {
					temp *= 2
				}
			}
		}
		totalSum += temp
	} 

	fmt.Println(totalSum)
}

type Card struct {
	win		map[int]int
	have	[]int
}

func DataParser()([]Card, error) {

	file, err := os.Open("./data.txt")
	data := []Card{}

	if err != nil {
		return data, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		card := strings.Split(line, ": ")[1]
		numberTuple := strings.Split(card, " | ")

		winningNumbers := strings.Split(numberTuple[0], " ")
		ourNumbers := strings.Split(numberTuple[1], " ")

		newCard := Card{
			win: ConvertListToMap[int](ConvertListStrToInt(winningNumbers)),
			have: ConvertListStrToInt(ourNumbers),
		}

		data = append(data, newCard)
	}

	return data, nil
}

func ConvertListStrToInt (list []string) []int {
	temp := []int{}
	for _, strNum := range list {
		if (strNum != "") {
			num, err := strconv.Atoi(strNum)

			if err != nil {
				fmt.Printf("error parsing: %s\n", strNum)
				break
			}

			temp =	append(temp, num)
		}
	}

	return temp
}

func ConvertListToMap[K int | string](list []K) (map[K]K) {
	result := map[K]K{}
	for _, value := range list {
		result[value] = value
	}
	return result
} 