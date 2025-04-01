package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-tapd/tapd/webhook"
)

type StoreUpdateListener struct{}

func (l *StoreUpdateListener) OnStoryUpdate(ctx context.Context, event *webhook.StoryUpdateEvent) error {
	log.Printf("StoreUpdateListener: %+v", event)
	return nil
}

func main() {
	dispatcher := webhook.NewDispatcher(
		webhook.WithRegisters(&StoreUpdateListener{}),
	)
	dispatcher.Registers(&StoreUpdateListener{})

	srv := http.NewServeMux()
	srv.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received webhook request")
		if err := dispatcher.DispatchRequest(r); err != nil {
			log.Println(err)
		}
		w.Write([]byte("ok")) //nolint:errcheck
	})

	http.ListenAndServe(":8080", srv) //nolint:errcheck
}
