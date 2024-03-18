package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	var greeting string
	switch {
	case currentTime.Hour() >= 7 && currentTime.Hour() < 12:
		greeting = "Goedemorgen!"
	case currentTime.Hour() >= 12 && currentTime.Hour() < 18:
		greeting = "Goedemiddag!"
	case currentTime.Hour() >= 18 && currentTime.Hour() < 23:
		greeting = "Goedenavond!"
	case currentTime.Hour() >= 23 && currentTime.Hour() < 7:
	}

	fmt.Println(greeting, "Welkom bij Fonteyn Vakantieparken.")
}
