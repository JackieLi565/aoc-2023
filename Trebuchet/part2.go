package main

import (
	"aoc-2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var WordToInt = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "4",
	"five":  "5e",
	"six":   "6",
	"seven": "7n",
	"eight": "e8t",
	"nine":  "9e",
}

func Trebuchet2() {	
	file, _ := os.Open("./data.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	temp := []int{}
	
	for scanner.Scan() {
		currentLine := scanner.Text()
		ReplaceStringToNumber(&currentLine)

		first, _ := FindFirstNumber(&currentLine, false);
		second, _ := FindFirstNumber(&currentLine, true);

		temp = append(temp, (first * 10) + second)
	} 

	fmt.Println(utils.Sum(&temp))
}

func ReplaceStringToNumber(str *string)  {
	for key, value := range WordToInt {
		*str = strings.ReplaceAll(*str, key, value)
	}
}