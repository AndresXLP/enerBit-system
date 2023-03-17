package server

import (
	"fmt"
	"net"

	"enerBit-system/config"
	pb "enerBit-system/internal/infra/grpc/service"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

type Server interface {
	Serve()
}

type Connection struct {
	protocol     string
	host         string
	port         int
	meterService pb.MeterServiceServer
}

func NewGrpcServer(server pb.MeterServiceServer) Server {
	return &Connection{
		protocol:     config.Environments().GrpcProtocol,
		host:         config.Environments().GrpcHost,
		port:         config.Environments().GrpcPort,
		meterService: server,
	}
}

func (c *Connection) Serve() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)

	listener, err := net.Listen(c.protocol, addr)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterMeterServiceServer(srv, c.meterService)

	fmt.Printf("⇋ gRPC server running on %s %s \n", c.protocol, color.Green(listener.Addr()))

	log.Fatal(srv.Serve(listener))
}
