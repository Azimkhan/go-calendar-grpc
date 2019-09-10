package main

import (
	"context"
	"fmt"
	calendarRpc "github.com/Azimkhan/go-calendar-grpc/internal/calendar/delivery/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calendarRpc.NewCalendarServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	response, err := c.ListEvents(ctx, &calendarRpc.ListCalendarEventRequest{})
	if err != nil {
		log.Fatalf("ListEvenets() error %v", err)
	}
	for _, event := range response.Events {
		fmt.Println(event)
	}

}
