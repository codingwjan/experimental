package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Event struct {
	Name string `json:"name"`
	Day  int    `json:"day"`
}

func restartpromt() {
	fmt.Println("do you want to restart the program? (y/n)")
	var restart string
	fmt.Scan(&restart)
	if restart == "y" {
		main()
	} else {
		return
	}

}

func main() {

	//look for file
	_, err := ioutil.ReadFile("events.json")

	file, err := ioutil.ReadFile("./events.json")

	if len(file) == 0 {
		fmt.Println("file is empty or does not exist")
		restartpromt()
	}

	var payload Event
	err = json.Unmarshal(file, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println(payload.Name)
	fmt.Println(payload.Day)

}