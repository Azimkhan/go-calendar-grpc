/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/Azimkhan/go-calendar-grpc/internal/configuration"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/repository"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/service"
	calendarGrpc "github.com/Azimkhan/go-calendar-grpc/internal/grpc"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var addr string

// grpcServerCmd represents the grpcServer command
var grpcServerCmd = &cobra.Command{
	Use:   "grpcServer",
	Short: "Run GRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := configuration.CreateLogger("configs/logger.json")
		logger.Info("Application started.")

		grpcServer := construct(logger)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve %v", err)
		}
	},
}

func init() {
	grpcServerCmd.Flags().StringVar(&addr, "address", "0.0.0.0:50051", "host:port to listen to")
}

func construct(logger *zap.Logger) *grpc.Server {
	grpcServer := grpc.NewServer()
	calendarRepo := repository.NewMapCalendarRepository()
	calendarService := service.NewCalendarService(calendarRepo, time.Second*15)
	server := calendarGrpc.NewCalendarServer(logger, calendarService)
	calendarGrpc.RegisterCalendarServiceServer(grpcServer, server)

	return grpcServer
}