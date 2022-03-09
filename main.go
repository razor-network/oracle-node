package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"razor/cmd"
	"time"
)

func main() {

	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 3

	err := sentry.Init(sentry.ClientOptions{
		Environment: "dev",
		Release:     "razor-go@0.1.76",
		Debug:       true,
		Transport:   sentrySyncTransport,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	defer sentry.Recover()

	cmd.InitializeInterfaces()
	cmd.Execute()
}
