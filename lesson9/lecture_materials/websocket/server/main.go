package main

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	addr = ":9000"
)

type wsService struct {
	connections map[int]net.Conn
	mu          sync.Mutex
	index       int
}

func (s *wsService) addConnection(conn net.Conn) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	i := s.index
	s.connections[s.index] = conn
	s.index += 1

	return i
}

func (s *wsService) chat(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("can't upgrade connection: %s\n", err.Error())
		return
	}

	connID := s.addConnection(conn)
	id := strconv.Itoa(connID)
	ch := make(chan []byte)

	go func() {
		defer func() {
			conn.Close()
			close(ch)
		}()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					log.Printf("can't read message from connection: %s\n", err.Error())
				}

				break
			}

			ch <- msg
		}
	}()

	go func() {
		for msg := range ch {
			msg = append([]byte(id+": "), msg...)
			s.mu.Lock()
			for key, connection := range s.connections {
				if err := wsutil.WriteServerMessage(connection, ws.OpText, msg); err != nil {
					log.Printf("can't write message: %s\n", err.Error())
					delete(s.connections, key)
				}
			}

			s.mu.Unlock()
		}
		log.Println("go func stop")
		s.mu.Lock()
		delete(s.connections, connID)
		s.mu.Unlock()
	}()
}

func main() {
	service := wsService{
		connections: make(map[int]net.Conn, 0),
		mu:          sync.Mutex{},
	}
	router := gin.Default()

	router.GET("/chat", gin.WrapF(service.chat))

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("can't listen and serve server: %s", err.Error())
	}
}
