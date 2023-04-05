package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"userservice/internal/adapters/pgrepo"
	"userservice/internal/app"
	"userservice/internal/ports/httpfiber"
)

func CreateDB(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, os.Getenv("DB_CONNECT_STRING"))
}

func main() {
	ctx := context.Background()

	dbConn, err := CreateDB(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := pgrepo.NewRepositoryPG(dbConn)

	usecase := app.NewApp(repo)

	server := httpfiber.NewHTTPServer(":3000", usecase)

	log.Fatal(server.Listen())
}
