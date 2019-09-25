package main

import (
	"github.com/Azimkhan/go-calendar-grpc/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
