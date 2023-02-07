package service

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	// internal GitHub

	// internal project
	pb "github.com/raythx98/go-database/pb"
	"github.com/raythx98/go-database/service/handler"
)

func StartService() {
	db, err := Init()
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	log.Println("Listening...")

	// can listen from any IP Address
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 9282))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterDatabaseServiceServer(grpcServer, &handler.Server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen to gRPC port 9282: %v", err)
	}
}
