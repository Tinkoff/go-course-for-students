package main

import (
	"context"
	"fmt"
	"gitlab.tcsbank.ru/a.krutyakov/lesson9/cmd/grpc/foo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.DialContext(context.Background(), "localhost:50054",
	//	grpc.WithChainUnaryInterceptor(ExampleUnaryClientInterceptor),
	//	grpc.WithChainStreamInterceptor(ExampleStreamClientInterceptor),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.DialContext(context.Background(), "localhost:50054",
	//	grpc.WithChainUnaryInterceptor(AuthUnaryClientInterceptor),
	//	grpc.WithChainStreamInterceptor(AuthStreamClientInterceptor),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	pricesClient := foo.NewPricesServiceClient(conn)

	lastPrice, err := pricesClient.LastPrice(context.Background(), &foo.Instrument{
		Id:     int64(123),
		Ticker: strconv.Itoa(123),
		Type:   foo.InstrumentType_Equity,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastPrice)

	clientStream, err := pricesClient.GetPrices(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		if err = clientStream.Send(&foo.Instrument{
			Id:     int64(i),
			Ticker: strconv.Itoa(i),
			Type:   foo.InstrumentType_Bond,
		}); err != nil {
			log.Fatal(err)
		}
	}

	prices, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nПолученный массив цен\n")
	for _, price := range prices.Prices {
		fmt.Println(price)
	}

	pricesStream, err := pricesClient.PricesStream(context.Background(), &foo.Instrument{
		Id:     int64(123),
		Ticker: strconv.Itoa(123),
		Type:   foo.InstrumentType_Equity,
		//AltName: "asd",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nПодключение к стримингку цен\n")
	for {
		price, err := pricesStream.Recv()
		if err == io.EOF {
			break
		}

		fmt.Println(price)
	}

	//chatClient := foo.NewChatServiceClient(conn)
	//
	//chat, err := chatClient.Chat(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//go func() {
	//	for {
	//		msg, err := chat.Recv()
	//		if err == io.EOF {
	//			break
	//		}
	//
	//		fmt.Println(msg.Payload)
	//	}
	//}()
	//
	//reader := bufio.NewReader(os.Stdin)
	//
	//for {
	//	fmt.Print("Enter text: ")
	//	text, err := reader.ReadString('\n')
	//	if err != nil {
	//		log.Fatalf("can't read string from stdin: %s", err.Error())
	//	}
	//
	//	if err := chat.Send(&foo.Message{Payload: text}); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	time.Sleep(100 * time.Millisecond)
	//}

	//barClient := foo.NewBarServiceClient(conn)
	//barMessage, err := barClient.Bar(context.Background(), &emptypb.Empty{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//switch data := barMessage.Data.(type) {
	//case *foo.BarMessage_Id:
	//	fmt.Println(data.Id)
	//case *foo.BarMessage_Value:
	//	fmt.Println(data.Value)
	//}
}

func ExampleUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	log.Println(method)

	return invoker(ctx, method, req, reply, cc, opts...)
}

type wrappedStream struct {
	grpc.ClientStream
}

func (s *wrappedStream) SendMsg(m interface{}) error {
	log.Println("send:", time.Now())

	return s.ClientStream.SendMsg(m)
}

func (s *wrappedStream) RecvMsg(m interface{}) error {
	log.Println("receive:", time.Now())

	return s.ClientStream.RecvMsg(m)
}

func ExampleStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("======= [Client Interceptor] ", method)

	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return &wrappedStream{s}, nil
}

const authToken = "112"

func AuthUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Basic: "+authToken)

	return invoker(ctx, method, req, reply, cc, opts...)
}

func AuthStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("======= [Client Interceptor] ", method)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Basic: "+authToken)

	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return s, nil
}
