package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"os"
	"time"
)

const url = "ws://localhost:9000/chat"

func main() {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), url)
	if err != nil {
		log.Fatalf("can't dial connection: %s", err.Error())
	}

	go func() {
		for {
			data, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				log.Printf("can't read server data: %s", err.Error())
				break
			}

			fmt.Printf("\n%s\n", data)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("can't read string from stdin: %s", err.Error())
		}

		if err := wsutil.WriteClientMessage(conn, ws.OpText, []byte(text)); err != nil {
			log.Fatalf("can't wtite client message: %s", err.Error())
		}

		time.Sleep(100 * time.Millisecond)
	}
}
