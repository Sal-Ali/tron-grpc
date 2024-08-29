package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Sal-Ali/tron-grpc/client"
	api "github.com/Sal-Ali/tron-grpc/tronapi"
	"google.golang.org/grpc"
)

func logData(v any) {
	s, err := json.Marshal(v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(s))
}

func main() {
	client := client.New("", "")
	err := client.Start(grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	var in api.EmptyMessage
	block, err := client.GetNowBlock2(context.Background(), &in)
	if err != nil {
		log.Fatalln(err)
	}

	logData(block)
}
