package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
)


func main(){
	listener, err:= net.Listen("tcp", ":8080")
	if err!=nil{
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Server is listening at PORT 8080")

	for{
		conn, err:= listener.Accept()
		if err !=nil{
			fmt.Println("Error accpeting connections", err)
			continue
		}
		fmt.Println("New Client connected!")

		go handleClient(conn)
	}
}

func handleClient( conn net.Conn){
	defer conn.Close()

	conn.Write([]byte("Welcome to the Go TCP Server!\n"))

	reader:=bufio.NewReader(conn)

	for{
		message, err:= reader.ReadString('\n')
		if err !=nil{
			fmt.Println("Client disconnected")
			return
		}

		fmt.Print("Message from Client", message)
		
		conn.Write([]byte("You said: " + message))

	}
}

