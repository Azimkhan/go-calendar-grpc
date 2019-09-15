package main

import (
	"context"
	"fmt"
	calendarRpc "github.com/Azimkhan/go-calendar-grpc/internal/calendar/delivery/grpc"
	"github.com/golang/protobuf/ptypes"
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

	client := calendarRpc.NewCalendarServiceClient(cc)

	// list events
	ListEvents(client)

	// create events
	now := time.Now()
	id1 := CreateEvent(now, "Clean the working desk", calendarRpc.EventType_TASK, err, client)
	_ = CreateEvent(now, "Meet with friends", calendarRpc.EventType_MEETING, err, client)

	// list events again
	ListEvents(client)

	// delete one events
	DeleteEvent(client, id1)

	// list events again
	ListEvents(client)
}

func DeleteEvent(client calendarRpc.CalendarServiceClient, id int64) {
	fmt.Printf("DeleteEvent(%d)\n", id)
	ctx := createContext()
	req := &calendarRpc.DeleteCalendarEventRequest{
		Id: id,
	}
	_, err := client.DeleteEvent(ctx, req)
	if err != nil {
		log.Fatalf("DeleteEvent() error %v", err)
	}
	fmt.Println("   Event deleted")
}

func CreateEvent(now time.Time, name string, eventType calendarRpc.EventType, err error, client calendarRpc.CalendarServiceClient) int64 {
	fmt.Println("CreateEvent()")
	ctx := createContext()
	start, _ := ptypes.TimestampProto(now)
	end, _ := ptypes.TimestampProto(now.Add(time.Hour))
	req2 := &calendarRpc.CreateCalendarEventRequest{
		Name:      name,
		Type:      eventType,
		StartTime: start,
		EndTime:   end,
	}
	response2, err := client.CreateEvent(ctx, req2)
	if err != nil {
		log.Fatalf("CreateEvent() error %v", err)
	}
	fmt.Printf("   Event created: %v\n", response2)
	return response2.Id
}

func ListEvents(client calendarRpc.CalendarServiceClient) {
	fmt.Println("ListEvents()")
	ctx := createContext()
	response1, err := client.ListEvents(ctx, &calendarRpc.ListCalendarEventRequest{})
	if err != nil {
		log.Fatalf("ListEvents() error %v", err)
	}
	PrintEvents(response1.Events)
}

func createContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 400*time.Millisecond)
	return ctx
}

func PrintEvents(events []*calendarRpc.CalendarEventObject) {
	for _, event := range events {
		fmt.Println("   ", event)
	}
}
