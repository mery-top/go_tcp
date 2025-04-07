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

	go func(){
		reader:= bufio.NewReader(conn)
		for{
			message,err:= reader.ReadString('\n')
			if err!=nil{
				fmt.Println("Client disconnected",err)
				return
			}
			fmt.Print("Client says:", message)
		}
	}()

	go func(){
		serverReader:= bufio.NewReader(os.Stdin)
		for{
			serverMessage, err:= serverReader.ReadString('\n')
			if err !=nil{
				fmt.Println("Server is not Connected", err)
				return
			}
			fmt.Print("Server Says:", serverMessage)

		}
	}()
}

