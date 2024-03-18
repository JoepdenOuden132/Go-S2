package main

import (
	"fmt"
	"strconv"
)

func main() {
	var interval int
	var input string
	var err error
	for {
		fmt.Println("How often does the alarm go off")
		fmt.Scanln(&input)
		interval, err = strconv.Atoi(input)

		if err != nil || interval <= 0 {
			fmt.Println("Wrong imput, try again.")
		} else {
			setAlarm(interval)
			break
		}
	}
}

func setAlarm(interval int) {

	fmt.Println("Alarm set")

	for i := 0; i < interval; i++ {
		fmt.Println("Alarm", (i + 1), "!")
	}

}
