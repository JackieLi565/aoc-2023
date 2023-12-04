package main

import (
	"aoc-2023/utils"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Trebuchet1() {
	file, _ := os.Open("./data.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	temp := []int{}
	
	for scanner.Scan() {
		currentLine := scanner.Text()

		first, _ := FindFirstNumber(&currentLine, false);
		second, _ := FindFirstNumber(&currentLine, true);

		temp = append(temp, (first * 10) + second)
	} 

	fmt.Println(utils.Sum(&temp))
}

func FindFirstNumber(arr *string, rev bool) (int, error) {
	length := len(*arr)
	if rev {
		for i := length - 1; i >= 0; i-- {
			isNumber, err := strconv.Atoi(string((*arr)[i]))

			if err != nil {
				continue
			} else {
				return isNumber, nil
			}
		}
	} else {
		for i := 0; i < length; i++ {
			isNumber, err := strconv.Atoi(string((*arr)[i]))

			if err != nil {
				continue
			} else {
				return isNumber, nil
			}
		}
	}

	return -1, errors.New("no valid number in array")
}