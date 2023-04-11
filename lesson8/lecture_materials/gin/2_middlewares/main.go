package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomMW(c *gin.Context) {
	t := time.Now()

	// c.Set("example", "12345")

	c.Next()

	latency := time.Since(t)
	status := c.Writer.Status()

	log.Println("latency", latency, "method", c.Request.Method, "path", c.Request.URL.Path, "status", status)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// r.Use(CustomMW)

	r.GET("/users/:id", GetUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func GetUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid name"})
		return
	}

	/*
		exampleData := c.GetString("example")
		if exampleData == "" {
			c.Status(http.StatusInternalServerError)
			return
		}
	*/

	c.Header("Server", "matrix")
	c.JSON(http.StatusOK,
		gin.H{
			"id":   c.Param("id"),
			"name": name,
			//"context": exampleData,
		},
	)
}
