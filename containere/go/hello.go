package main

import (
	"fmt"
	"strconv"
	"strings"
)

var daysOfMonth = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(dateToDayNumber("2019-01-09")) // expect 9
	fmt.Println(dateToDayNumber("2019-02-10")) // expect 41
	fmt.Println(dateToDayNumber("2016-03-02")) // expect 62
}

func stringToInt(str string) int {
	var i, err = strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func dateToDayNumber(date string) int {
	splitDate := strings.Split(date, "-")
	year := stringToInt(splitDate[0])
	month := stringToInt(splitDate[1])
	day := stringToInt(splitDate[2])
	result := 0
	for i := 0; i < month - 1; i++ {
		result += daysOfMonth[i]
	}
	if (year % 4 == 0 && month > 2) {
		result++
	}
	return result + day
}
