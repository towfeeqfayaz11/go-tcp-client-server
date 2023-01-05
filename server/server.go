package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	fmt.Println("Message received from client is: ", string(buf[:n]))
	conn.Write([]byte("Hello client, Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func main() {
	// addr, err := net.ResolveTCPAddr(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	// if err != nil {
	// 	panic(err)
	// }
	// // Listener for incoming connections.
	// l, err := net.ListenTCP(CONN_TYPE, addr)

	// Listener for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("The TCP Server Will Start Listening on " + CONN_HOST + ":" + CONN_PORT)
	var i = 1
	for {
		// Listen for an incoming connection.
		fmt.Println("listening for client")
		conn, err := l.Accept()
		// conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// err = conn.SetKeepAlive(true)
		fmt.Println("client", i, "connected")
		i++
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// If you build this and run it, you'll have a simple TCP server running on port 3333.
// To test your server, send some raw data to that port:

// echo -n "test out the server" | nc localhost 3333
// You should get a response:

// "Message received."

// or

// curl telnet://localhost:3333
