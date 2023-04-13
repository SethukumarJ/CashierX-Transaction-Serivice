package http

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"

	pb "github.com/SethukumarJ/CashierX-Auth-Service/pkg/pb"
	"google.golang.org/grpc"

	_ "github.com/SethukumarJ/CashierX-Auth-Service/cmd/api/docs"
	handler "github.com/SethukumarJ/CashierX-Auth-Service/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func StartGRPCServer(accountHandler *handler.AccountHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, accountHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(accountHandler *handler.AccountHandler) *ServerHTTP {
	engine := gin.New()
	go StartGRPCServer(accountHandler, "50057")

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3003")
}
