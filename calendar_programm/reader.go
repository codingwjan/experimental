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


//get the events from the file
func getEvents() []Event {
	//read the file
	file, err := ioutil.ReadFile("events.json")
	if err != nil {
		log.Fatal(err)
	}

	//convert the file to a slice of events
	var events []Event
	err = json.Unmarshal(file, &events)
	if err != nil {
		log.Fatal(err)
	}
	return events
}




func main() {
	getEvents()
	fmt.Println("welcome to the calendar program")
	fmt.Println("you are currently in the reader mode")
	fmt.Println("here are your events")
	for _, event := range getEvents() {
		fmt.Println(event.Name, event.Day)
	}

}