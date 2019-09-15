package main

import (
	calendarGrpc "github.com/Azimkhan/go-calendar-grpc/internal/calendar/delivery/grpc"
	"github.com/Azimkhan/go-calendar-grpc/internal/calendar/repository"
	"github.com/Azimkhan/go-calendar-grpc/internal/calendar/usecase"
	"github.com/Azimkhan/go-calendar-grpc/internal/configuration"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
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

	calendarRepo := repository.NewMapCalendarRepository()
	calendarUsecase := usecase.NewCalendarUsecase(calendarRepo, time.Second*15)
	server := calendarGrpc.NewCalendarServer(logger, calendarUsecase)
	calendarGrpc.RegisterCalendarServiceServer(grpcServer, server)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
