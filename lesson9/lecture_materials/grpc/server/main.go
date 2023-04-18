package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitlab.tcsbank.ru/a.krutyakov/lesson9/cmd/grpc/foo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	port := ":50054"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := &ChatService{
		connections: make(map[int]foo.ChatService_ChatServer),
	}

	//server := grpc.NewServer()
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(ExampleUnaryServerInterceptor), grpc.ChainStreamInterceptor(ExampleStreamInterceptor))
	//server := grpc.NewServer(grpc.ChainUnaryInterceptor(AuthUnaryServerInterceptor), grpc.ChainStreamInterceptor(AuthStreamInterceptor))
	foo.RegisterPricesServiceServer(server, &PricesService{})
	foo.RegisterChatServiceServer(server, service)
	//foo.RegisterBarServiceServer(server, &BarService{})

	log.Printf("Starting gRPC listener on port " + port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type PricesService struct{}

func (s *PricesService) LastPrice(ctx context.Context, instrument *foo.Instrument) (*foo.Price, error) {
	fmt.Println(instrument)

	return &foo.Price{
		InstrumentID: instrument.Id,
		Value:        rand.Float64() * 100,
		Ts:           timestamppb.New(time.Now().UTC()),
	}, nil
}

func (s *PricesService) GetPrices(stream foo.PricesService_GetPricesServer) error {
	out := make([]*foo.Price, 0)
	for {
		instrument, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&foo.Prices{Prices: out})
		}

		fmt.Println(instrument)

		out = append(out, &foo.Price{
			InstrumentID: instrument.Id + 1,
			Value:        rand.Float64() * 100,
			Ts:           timestamppb.New(time.Now().UTC()),
		})
	}
}

func (s *PricesService) PricesStream(instrument *foo.Instrument, stream foo.PricesService_PricesStreamServer) error {
	fmt.Println(instrument)

	for i := 0; i < 5; i++ {
		if err := stream.Send(&foo.Price{
			InstrumentID: instrument.Id + 1,
			Value:        rand.Float64() * 100,
			Ts:           timestamppb.New(time.Now().UTC()),
		}); err != nil {
			return err
		}
	}
	return nil
}

type ChatService struct {
	connections map[int]foo.ChatService_ChatServer
	mu          sync.Mutex
	index       int
}

func (s *ChatService) addConnection(stream foo.ChatService_ChatServer) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	i := s.index
	s.connections[s.index] = stream
	s.index += 1

	return i
}

func (s *ChatService) Chat(stream foo.ChatService_ChatServer) error {
	connID := s.addConnection(stream)
	id := strconv.Itoa(connID)
	ch := make(chan string)

	go func() {
		defer func() {
			close(ch)
		}()

		for {
			msg, err := stream.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					log.Printf("can't read message from connection: %s\n", err.Error())
				}

				break
			}

			ch <- msg.Payload
		}
	}()

	for msg := range ch {
		msg = fmt.Sprintf("%s: %s", id, msg)
		s.mu.Lock()
		for key, otherStream := range s.connections {
			if err := otherStream.Send(&foo.Message{Payload: msg}); err != nil {
				log.Printf("can't write message: %s\n", err.Error())
				delete(s.connections, key)
			}
		}

		s.mu.Unlock()
	}

	return nil
}

func ExampleUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod)

	return handler(ctx, req)
}

type wrappedStream struct {
	grpc.ServerStream
}

func (s *wrappedStream) SendMsg(m interface{}) error {
	log.Println("send:", time.Now())

	return s.ServerStream.SendMsg(m)
}

func (s *wrappedStream) RecvMsg(m interface{}) error {
	log.Println("receive:", time.Now())

	return s.ServerStream.RecvMsg(m)
}

func ExampleStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println(info.FullMethod)

	return handler(srv, &wrappedStream{ss})
}

const authToken = "112"

func AuthUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Metadata not found")
	}

	authStr, ok := md["authorization"]
	if !ok || len(authStr) == 0 {
		return nil, status.Error(codes.Unauthenticated, "no auth details")
	}

	authArr := strings.Split(authStr[0], " ")
	if len(authArr) != 2 {
		return nil, status.Error(codes.Unauthenticated, "bad auth data")
	}

	if authArr[1] != authToken {
		return nil, status.Error(codes.Unauthenticated, "incorrect token")
	}

	return handler(ctx, req)
}

func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Error(codes.InvalidArgument, "retrieving metadata failed")
	}

	authStr, ok := md["authorization"]
	if !ok || len(authStr) == 0 {
		return status.Error(codes.Unauthenticated, "no auth details")
	}

	authArr := strings.Split(authStr[0], " ")
	if len(authArr) != 2 {
		return status.Error(codes.Unauthenticated, "bad auth data")
	}

	if authArr[1] != authToken {
		return status.Error(codes.Unauthenticated, "incorrect token")
	}

	return handler(srv, ss)
}

type BarService struct{}

func (s *BarService) Bar(ctx context.Context, empty *emptypb.Empty) (*foo.BarMessage, error) {
	return &foo.BarMessage{
		Data: &foo.BarMessage_Id{Id: "data"},
	}, nil
}
