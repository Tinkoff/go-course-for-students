package main

import (
	"log"
	"net/http"

	_ "lesson8/swagger/docs"

	swagger "github.com/swaggo/http-swagger"
)

type myError struct {
	Status int
	Error  string
}

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get user by ID
//	@ID				get-user-by-int
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	model.Error
//	@Failure		404	{object}	model.Error
//	@Failure		500	{object}	myError
//	@Router			/user/{id} [get]
func handleUsers(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}

//	@title			Sample Project API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		petstore.swagger.io
// @BasePath	/api/v1
func main() {
	http.HandleFunc("/docs/", swagger.WrapHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello, " + r.Header.Get("User-Agent") + ", with url " + r.URL.String()))
	})

	log.Println("server starts")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
