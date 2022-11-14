package main

import (
	"fmt"
	"bufio"
	"os"
	//"strconv"
	//"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first name: ")
	firstName, _ := reader.ReadString('\n')
	fmt.Print(firstName)
}
