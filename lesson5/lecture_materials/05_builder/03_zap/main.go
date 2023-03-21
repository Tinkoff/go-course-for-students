package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("can;t init logger: %v", err))
	}

	defer logger.Sync()

	logger.
		WithOptions(zap.WithCaller(true)).
		With(zap.String("key", "val")).
		Error("hello! i'm error!")
}
