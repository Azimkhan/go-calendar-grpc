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
	"context"
	"fmt"
	calendarRpc "github.com/Azimkhan/go-calendar-grpc/internal/grpc"
	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

var target string

// grpcClientCmd represents the grpcClient command
var grpcClientCmd = &cobra.Command{
	Use:   "grpcClient",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cc, err := grpc.Dial(target, grpc.WithInsecure())
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
	},
}

func init() {
	grpcClientCmd.Flags().StringVar(&target, "target", "localhost:50051", "target host:port")
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func PrintEvents(events []*calendarRpc.CalendarEventObject) {
	for _, event := range events {
		fmt.Println("   ", event)
	}
}
