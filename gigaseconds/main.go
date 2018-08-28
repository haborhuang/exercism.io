package main

import (
	"log"
	"time"
)

func main() {
	log.Println(getMoment(time.Now()))
}

func getMoment(birthday time.Time) time.Time {
	return birthday.Add(1e9 * time.Second)
}
