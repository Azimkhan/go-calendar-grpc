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
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/configuration"
	"github.com/Azimkhan/go-calendar-grpc/internal/maindb"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
)

// Starts worker process that produces reminder
var WorkCmd = &cobra.Command{
	Use:   "work",
	Short: "Monitors calendar and produces reminders",
	Run: func(cmd *cobra.Command, args []string) {
		logger := configuration.CreateLogger("configs/logger.json")
		logger.Info("Reminder worker started.")
		repo, err := maindb.NewPgEventRepository(dsn)

		if err != nil {
			logger.Fatal("Unable to create repository", zap.Error(err))
		}

		work := func(from time.Time, to time.Time, ctx context.Context) {
			events, err := repo.FetchByDateRange(from, to, ctx)
			if err != nil {
				logger.Warn("Worker fetch error", zap.Error(err))
				return
			}

			for _, e := range events {
				logger.Info("event reminder", zap.Any("event", e))
			}
		}

		for {
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			from := time.Now().UTC()
			to := from.Add(1 * time.Minute)
			work(from, to, ctx)
			dur := to.Sub(time.Now())
			time.Sleep(dur)
		}

	},
}

var dsn string

func init() {
	WorkCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=calendar password=calendar dbname=calendar", "database connection string")
}
