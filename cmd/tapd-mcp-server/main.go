package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-tapd/tapd"
	"github.com/go-tapd/tapd/mcp"
)

func init() {
	requiredEnvs("TAPD_CLIENT_ID", "TAPD_CLIENT_SECRET", "TAPD_WORKSPACE_ID")
}

func main() {
	var (
		clientID     = os.Getenv("TAPD_CLIENT_ID")
		clientSecret = os.Getenv("TAPD_CLIENT_SECRET")
		workspace    = os.Getenv("TAPD_WORKSPACE_ID")
	)

	if clientID == "" || clientSecret == "" || workspace == "" {
		log.Fatal("missing TAPD_CLIENT_ID, TAPD_CLIENT_SECRET or TAPD_WORKSPACE_ID")
	}

	workspaceID, err := convertToInt(workspace)
	if err != nil {
		log.Fatalf("invalid TAPD_WORKSPACE_ID: %s", err)
	}

	client, err := tapd.NewClient(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := mcp.NewServer(workspaceID, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.ServerStdio(); err != nil {
		log.Fatal(err)
	}
}

func requiredEnvs(keys ...string) {
	var missing []string
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("missing required env vars: %v", missing)
	}
}

func convertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid int: %s", s)
	}
	return i, nil
}
