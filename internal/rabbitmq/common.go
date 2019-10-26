package rabbitmq

import "go.uber.org/zap"

func failOnError(err error, msg string, logger *zap.Logger) {
	if err != nil {
		logger.Fatal(msg, zap.Error(err))
	}
}
