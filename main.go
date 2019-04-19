package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/francescomalatesta/loggo/commands"
)

func main() {
	cm := map[string]commands.Command{
		"hello": commands.HelloCommand{},
	}

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(1)
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Loggo Server Started at localhost:" + PORT)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, cm)
	}
}

func handleConnection(c net.Conn, cm map[string]commands.Command) {
	fmt.Printf("CONNECTED - %s\n", c.RemoteAddr().String())

	mustClose := false
	for !mustClose {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		textCommand := strings.TrimSpace(string(netData))
		if textCommand == "EXIT" {
			break
		}

		r, err := HandleInputCommandAndReturnResult(textCommand, cm)
		if err != nil {
			r = err.Error()
		}

		c.Write([]byte(r + "\n"))
	}

	c.Close()

	fmt.Printf("CLOSED - %s\n", c.RemoteAddr().String())
}
