package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	//"os"
)


func main() {
	//loop forever
	
		//connect to the server
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		//send the message
		fmt.Fprintf(conn, "hello world\n")
		//read the response
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message from server: " + message)
		//close the connection
		conn.Close()
}

