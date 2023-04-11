package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string
	Password string
	Info     string
}

type Storage struct {
	mx   *sync.RWMutex
	data map[string]*User
}

func (s *Storage) Get(id string) (User, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	user, ok := s.data[id]
	if !ok || user == nil {
		return User{}, false
	}

	return *user, true
}

func (s *Storage) Set(id string, user User) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.data[id] = &user
}

var storage = Storage{
	mx: &sync.RWMutex{},
	data: map[string]*User{
		"0": {
			Name:     "Groot",
			Password: "qwe123",
			Info:     "I am Groot",
		},
	},
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusNotFound, "text/plain; charset=utf-8", []byte("hello"))
		// c.String(http.StatusNotFound, "hello")
	})

	r.GET("/users/:id", GetUser)
	r.PUT("/users/:id", UpdateUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func GetUser(c *gin.Context) {
	c.Header("Server", "matrix")

	id := c.Param("id")
	user, ok := storage.Get(id)
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	c.Header("Server", "matrix")

	id := c.Param("id")
	user, ok := storage.Get(id)
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}

	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name"})
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user.Name = name
	user.Info = string(body)
	storage.Set(id, user)
}
