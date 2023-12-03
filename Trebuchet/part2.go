package main

import (
	"aoc-2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	fmt.Println(Sum(&temp))
}

func ReplaceStringToNumber(str *string)  {
	for key, value := range utils.WordToInt {
		*str = strings.ReplaceAll(*str, key, value)
	}
}