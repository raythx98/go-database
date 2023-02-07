package handler

import (
	"fmt"
	database_proto "github.com/raythx98/go-database/pb"
	"log"

	"github.com/raythx98/go-database/service/command"
	"github.com/raythx98/go-database/service/invoker"
	"github.com/raythx98/go-database/service/receiver"
	"golang.org/x/net/context"
)

type Server struct {
	database_proto.UnimplementedDatabaseServiceServer
}

func (s *Server) AddLink(ctx context.Context, request *database_proto.AddLinkRequest) (*database_proto.AddLinkResponse, error) {
	addLinkCommand := &command.AddLinkCommand{
		Database: receiver.DbInstance,
		FullLink: request.GetFullLink(),
	}

	if customLink := request.GetCustomLink(); customLink != GrpcDefaultString {
		addLinkCommand.CustomLink = customLink
	}
	if invalidateAt := request.GetInvalidateAt(); invalidateAt != GrpcDefaultString {
		addLinkCommand.InvalidateAt = invalidateAt
	}
	if numRedirects := request.GetNumRedirects(); numRedirects != GrpcDefaultInt {
		addLinkCommand.NumRedirects = int(numRedirects)
	}

	addLinkInvoker := &invoker.Invoker{
		Command: addLinkCommand,
	}
	shortenedUrl, err := addLinkInvoker.Invoke()
	log.Println(fmt.Sprintf("grpc response:%s for request {%+v}", shortenedUrl, request))

	resp := &database_proto.AddLinkResponse{
		ShortenedUrl: shortenedUrl,
	}
	return resp, err
}

func (s *Server) GetFullLink(ctx context.Context, request *database_proto.GetFullLinkRequest) (*database_proto.GetFullLinkResponse, error) {
	getFullLinkCommand := &command.GetFullLinkCommand{
		Database:      receiver.DbInstance,
		ShortenedLink: request.GetShortenedUrl(),
	}

	getFullLinkInvoker := &invoker.Invoker{
		Command: getFullLinkCommand,
	}
	fullLink, err := getFullLinkInvoker.Invoke()
	log.Println(fmt.Sprintf("grpc response:%s for request {%+v}", fullLink, request))

	resp := &database_proto.GetFullLinkResponse{
		FullLink: fullLink,
	}
	return resp, err
}
