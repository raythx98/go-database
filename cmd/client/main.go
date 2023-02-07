package main

import (
	pb "github.com/raythx98/go-database/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	conn     *grpc.ClientConn
	grpcConn pb.DatabaseServiceClient
)

func main() {
	initGrpc()
	callAddLink()
	callGetFullLink()
	closeGrpc()
}

func initGrpc() {
	conn, err := grpc.Dial(":9282", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	grpcConn = pb.NewDatabaseServiceClient(conn)
}

func closeGrpc() {
	if conn == nil {
		return
	}
	err := conn.Close()
	if err != nil {
		log.Fatalf("Failed to close grpc connection %v", err)
	}
}

func callAddLink() {
	log.Println()
	addLinkRequest := &pb.AddLinkRequest{
		FullLink: "www.google.com",
	}
	addLinkResponse, err := grpcConn.AddLink(context.Background(), addLinkRequest)
	if err != nil {
		log.Fatalf("Error when calling AddLink: %s", err)
	}
	log.Printf("Request:{%+v} Response:{%+v}\n", addLinkRequest, addLinkResponse)
}

func callGetFullLink() {
	getFullLinkRequest := &pb.GetFullLinkRequest{
		ShortenedUrl: "eFx9lsVB",
	}
	getFullLinkResponse, err := grpcConn.GetFullLink(context.Background(), getFullLinkRequest)
	if err != nil {
		log.Fatalf("Error when calling GetFullLink: %s", err)
	}
	log.Printf("Request:{%+v} Response:{%+v}\n", getFullLinkRequest, getFullLinkResponse)
}
