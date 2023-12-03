package main

import (
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

	fmt.Println(Sum(&temp))
}

func Sum(arr *[]int) int {
	total := 0

	for _, num := range *arr {
		total += num
	}

	return total
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