package initiable

// import (
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"
// 	"word-book/config"
// )

// var Logger zerolog.Logger

// func initSentry() {
// 	dsn := config.C.Sentry.Dsn
// 	debug, err := strconv.ParseBool(config.C.RunMode)
// 	if err != nil {
// 		log.Fatalf("debug等级配置异常: %s", err)
// 	}
// 	err = sentry.Init(sentry.ClientOptions{
// 		// Either set your DSN here or set the SENTRY_DSN environment variable.
// 		Dsn: dsn,
// 		// Either set environment and release here or set the SENTRY_ENVIRONMENT
// 		// and SENTRY_RELEASE environment variables.
// 		Environment: "",
// 		Release:     "",
// 		// Enable printing of SDK debug messages.
// 		// Useful when getting started or trying to figure something out.
// 		Debug:            debug,
// 		AttachStacktrace: true,
// 	})
// 	if err != nil {
// 		log.Fatalf("sentry.Init Failed: %s", err)
// 	}
// 	// Flush buffered events before the program terminates.
// 	// Set the timeout to the maximum duration the program can afford to wait.
// 	defer sentry.Flush(2 * time.Second)
// }

// type sendSentry struct {
// }

// func (l *sendSentry) Run(e *zerolog.Event, level zerolog.Level, message string) {
// 	confLevel := config.C.Sentry.Level
// 	if zerolog.Level(confLevel) <= level {
// 		sentry.CaptureMessage(message)
// 	}

// }

// func InitLog() zerolog.Logger {
// 	initSentry()
// 	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
// 	Logger = logger.Hook(&sendSentry{})
// 	return Logger
// }
