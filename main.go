package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"micro-grcp/client"
	"micro-grcp/proto"
	"time"
)

func main() {
	var (
		json = flag.String("listenaddr", ":3000", "listen the address the service is running")
		grpc = flag.String("grpc", ":4000", "listen the address of the grpc transport")
		ctx  = context.Background()
		svc  = NewLoggingService(NewMetricService(&priceFetcher{}))
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal()
	}

	go func() {
		for {
			coins := []string{"BTC", "ETH", "GG"}
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: coins[rand.Intn(len(coins))]})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}

	}()

	go makeGRPCServerAndRun(*grpc, svc)

	jsonServer := NewJSONAPIServer(*json, svc)
	jsonServer.Run()

}
