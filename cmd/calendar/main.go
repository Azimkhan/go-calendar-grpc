package main

import (
	calendarGrpc "github.com/Azimkhan/go-calendar-grpc/internal/calendar/delivery/grpc"
	"github.com/Azimkhan/go-calendar-grpc/internal/configuration"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger := configuration.CreateLogger("configs/logger.json")
	logger.Info("Application started.")
	// run service code goes here...

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	calendarGrpc.RegisterCalendarServiceServer(grpcServer, &calendarGrpc.CalendarServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
