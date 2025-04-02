# 🚀 Go-Tapd-SDK

![Supported Go Versions](https://img.shields.io/badge/Go-%3E%3D1.23-blue)
[![Package Version](https://badgen.net/github/release/go-tapd/tapd/stable)](https://github.com/go-tapd/tapd/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-tapd/tapd)](https://pkg.go.dev/github.com/go-tapd/tapd)
[![codecov](https://codecov.io/gh/go-tapd/tapd/graph/badge.svg?token=QPTHZ5L9GT)](https://codecov.io/gh/go-tapd/tapd)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-tapd/tapd)](https://goreportcard.com/report/github.com/go-tapd/tapd)
[![lint](https://github.com/go-tapd/tapd/actions/workflows/lint.yml/badge.svg)](https://github.com/go-tapd/tapd/actions/workflows/lint.yml)
[![tests](https://github.com/go-tapd/tapd/actions/workflows/test.yml/badge.svg)](https://github.com/go-tapd/tapd/actions/workflows/test.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

The Go-Tapd-SDK is a Go client library for accessing the [Tapd API](https://www.tapd.cn/).

> [!WARNING]  
> This is currently still a non-stable version, is not recommended for production use. 

If you encounter any issues, you are welcome to [submit an issue](https://github.com/go-tapd/tapd/issues/new).

## 📥 Installation

```bash
go get github.com/go-tapd/tapd
```

## ✨ Features

see [features.md](features.md)

## 🔧 Usage

### API Service

```go
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
		WorkspaceID: tapd.Ptr(123456),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("labels: %+v", labels)
}
```

### Webhook Server Example

```go
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
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", srv)
}
```

## 📜 License

The MIT License (MIT). Please see [License File](LICENSE) for more information.