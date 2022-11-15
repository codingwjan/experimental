package main

import (
	"net"
	"log"
	
)

//handle the connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	//read from the connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}
	//print the message
	log.Println(string(buffer[:n]))

	//if the message is "hello" send a response
	if string(buffer[:n]) == "hello world\n" {
		conn.Write([]byte("hello client\n"))
	}
}

func main() {
	//create a listener
	log.Printf("starting server on port 8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	//accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}

}



