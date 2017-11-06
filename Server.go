package main

import "net"
import "fmt"
import "bufio"
import (

	//"time"
	"strings"

) // only needed below for sample processing


func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	ln1, _ := net.Listen("tcp", ":8082")

	// accept connection on port
	conn, _ := ln.Accept()
	conn1, _ := ln1.Accept()

	// run loop forever (or until ctrl-c)

	for {

		// will listen for message to process ending in newline (\n)
		//start1 := time.Now()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//start2 := time.Now()
		message1, _ := bufio.NewReader(conn1).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		fmt.Print("Another Message Received:", string(message1))



		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
		//t1 := time.Now()
		//elapsed1 := t1.Sub(start1)
		//fmt.Print("Message Received:", elapsed1)

		newmessage1 := strings.ToUpper(message1)
		// send new string back to client
		conn1.Write([]byte(newmessage1 + "\n"))
		//t2 := time.Now()
		//elapsed2 := t2.Sub(start2)
		//fmt.Print("Message Received:", elapsed2)
		if message == message1  {
			return_message := "Sorry, this ticket is already taken! Please introduce a new one!"
			conn1.Write([]byte(return_message + "\n"))
			newmessage1 := strings.ToUpper(message1)
			// send new string back to client
			conn1.Write([]byte(newmessage1 + "\n"))		} /*else {
			return_message := "Sorry, this ticket is already taken! Please introduce a new one!"
			conn.Write([]byte(return_message + "\n"))
		}*/

	}
}
