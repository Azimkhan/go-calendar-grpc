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
package reminder

import (
	"github.com/Azimkhan/go-calendar-grpc/internal/configuration"
	"github.com/Azimkhan/go-calendar-grpc/internal/rabbitmq"
	"github.com/spf13/cobra"
)

// Starts process that consumes reminders
var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Shows incoming reminders",
	Run: func(cmd *cobra.Command, args []string) {
		logger := configuration.CreateLogger("configs/logger.json")
		logger.Info("Reminder consumer started.")

		consumer := rabbitmq.NewConsumer(clientAmqpUrl, workerQueueName, logger)
		consumer.Run()
	},
}

var clientAmqpUrl string
var clientAmqpQueue string

func init() {
	ConsumerCmd.Flags().StringVar(&clientAmqpUrl, "amqpUrl", "amqp://guest:guest@localhost:5672/", "AMQP connection string")
	ConsumerCmd.Flags().StringVar(&clientAmqpQueue, "amqpQueue", "reminders", "AMQP queue name")
}
