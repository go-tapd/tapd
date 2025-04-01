package main

import (
	"context"
	"log"

	"github.com/go-tapd/tapd"
)

func main() {
	client, err := tapd.NewClient("client_id", "client_secret")
	if err != nil {
		log.Fatal(err)
	}

	// example: get labels
	labels, _, err := client.LabelService.GetLabels(context.Background(), &tapd.GetLabelsRequest{
		WorkspaceID: tapd.Ptr(123456), // nolint:mnd
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("labels: %+v", labels)
}
