package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	cookieName  = "session_id"
	cookieValue = "ID"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/sessions", Login)

	// r.DELETE("/sessions", CheckAuth, Login) // todo: logout

	authorized := r.Group("/users", CheckAuth)
	{
		authorized.GET("/:id", GetUser)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func CheckAuth(c *gin.Context) {
	value, err := c.Cookie(cookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if value != cookieValue {
		/*
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    cookieName,
				Value:   cookieValue,
				Expires: time.Now().Add(-time.Hour),
				Path:    "/",
			})
		*/
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

func Login(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    cookieName,
		Value:   cookieValue,
		Expires: time.Now().Add(time.Hour),
		Path:    "/",
	})

	c.JSON(http.StatusOK,
		gin.H{
			"name": c.DefaultQuery("name", "guest"),
		},
	)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"id":   c.Param("id"),
			"name": c.DefaultQuery("name", "guest"),
		},
	)
}
