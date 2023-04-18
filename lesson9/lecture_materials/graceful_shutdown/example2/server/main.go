package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.tcsbank.ru/a.krutyakov/lesson9/cmd/grpc/foo"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = ":50054"
	httpPort = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcService := &GRPCService{}
	grpcServer := grpc.NewServer()
	foo.RegisterBarServiceServer(grpcServer, grpcService)

	httpService := &HTTPService{}
	router := gin.Default()
	router.GET("/foo", gin.WrapF(httpService.foo))

	httpServer := http.Server{
		Addr:    httpPort,
		Handler: router,
	}

	eg, ctx := errgroup.WithContext(context.Background())

	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	// run grpc server
	eg.Go(func() error {
		log.Printf("starting grpc server, listening on %s\n", grpcPort)
		defer log.Printf("close grpc server listening on %s\n", grpcPort)

		errCh := make(chan error)

		defer func() {
			grpcServer.GracefulStop()
			_ = lis.Close()

			close(errCh)
		}()

		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("grpc server can't listen and serve requests: %w", err)
		}
	})

	eg.Go(func() error {
		log.Printf("starting http server, listening on %s\n", httpServer.Addr)
		defer log.Printf("close http server listening on %s\n", httpServer.Addr)

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s: %s", httpServer.Addr, err.Error())
			}

			close(errCh)
		}()

		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	})

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}

	log.Println("servers were successfully shutdown")
}

type GRPCService struct{}

func (s *GRPCService) Bar(ctx context.Context, _ *emptypb.Empty) (*foo.BarMessage, error) {
	log.Println("GRPCService: get request")
	time.Sleep(10 * time.Second)
	log.Println("GRPCService: write response")
	return &foo.BarMessage{
		Data: &foo.BarMessage_Id{Id: "bar"},
	}, nil
}

type HTTPService struct{}

func (s *HTTPService) foo(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTPService: get request")
	time.Sleep(10 * time.Second)
	log.Println("HTTPService: write response")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("foo")); err != nil {
		log.Println(err)
	}
}

//func (s *GRPCService) ClientStream(_ foo.FooBarService_ClientStreamServer) error { return nil }
//
//func (s *GRPCService) ServerStream(_ *foo.ObjectID, _ foo.FooBarService_ServerStreamServer) error {
//	return nil
//}
//
//func (s *GRPCService) Chat(_ foo.FooBarService_ChatServer) error { return nil }
