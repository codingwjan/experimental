package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

type Event struct {
	Name string `json:"name"`
	Day int `json:"day"`

}

func main() {

	//look for file
	_, err := os.Stat("events.json")
	if err != nil {
		//create file
		file, err := os.Create("events.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	var day int

	fmt.Println("welcome to the calendar program")
	fmt.Println("you are currently in the creator mode")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the name of the event")
	name, _ := reader.ReadString('\n')

	fmt.Println("please enter the date of the event")
	fmt.Scan(&day)

	Event := Event{Name: name, Day: day}

	byteArray, err := json.Marshal(Event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))	


	//write to json file
	file, err := os.OpenFile("events.json", os.O_APPEND|os.O_WRONLY, 0600)
	file.Write(byteArray)
	file.Sync()


	


}

