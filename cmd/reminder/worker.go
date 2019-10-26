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
	"github.com/Azimkhan/go-calendar-grpc/internal/maindb"
	"github.com/Azimkhan/go-calendar-grpc/internal/rabbitmq"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Starts worker process that produces reminder
var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Monitors calendar and produces reminders",
	Run: func(cmd *cobra.Command, args []string) {
		logger := configuration.CreateLogger("configs/logger.json")
		repo, err := maindb.NewPgEventRepository(dsn)

		if err != nil {
			logger.Fatal("Unable to create repository", zap.Error(err))
		}

		worker := rabbitmq.NewWorker(workerAmqpUrl, workerQueueName, repo, logger)
		defer worker.Close()

		worker.Run()
	},
}

var dsn string
var workerAmqpUrl string
var workerQueueName string

func init() {
	WorkerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=calendar password=calendar dbname=calendar", "database connection string")
	WorkerCmd.Flags().StringVar(&workerAmqpUrl, "amqpUrl", "amqp://guest:guest@localhost:5672/", "AMQP connection string")
	WorkerCmd.Flags().StringVar(&workerQueueName, "amqpQueue", "reminders", "AMQP queue name")
}
