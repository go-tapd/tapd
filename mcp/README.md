# Tapd MCP Server

## 📥 Usage

### Use STDIO Server

**Build the tapd-mcp-server**

```bash
git clone git@github.com:go-tapd/mcp.git

cd mcp && make build/cmd/tapd-mcp-server

./bin/tapd-mcp-server # <--- This is the command to run the STDIO server.
```

**Configure the MCP server**

Below is a configuration example based on Cline, with different configurations for various MCP Clients.

```json
{
  "mcpServers": {
    "github.com/go-tapd/mcp": {
      "command": "{path}/tapd-mcp-server",
      "env": {
        "TAPD_CLIENT_ID": "<YOUR_CLIENT_ID>",
        "TAPD_CLIENT_SECRET": "<YOUR_CLIENT_SECRET>",
        "TAPD_WORKSPACE_ID": "<YOUR_WORKSPACE_ID>"
      }
    }
  }
}
```

### Use SSE Server

**Install the package**

```bash
go get github.com/go-tapd/mcp
```

**Create a server**

```go
package main

import (
	"log"
	"net/http"

	"github.com/go-tapd/mcp"
	"github.com/go-tapd/tapd"
)

func main() {
	client, err := tapd.NewClient("client_id", "client_secret")
	if err != nil {
		log.Fatal(err)
	}

	workspaceID := 123456 // replace with your workspace ID

	srv, err := mcp.NewServer(workspaceID, client)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", srv.ServeHTTP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

Visit http://localhost:8080/sse to get the SSE stream.

## 📦 Features

### 需求

- [x] [返回符合查询条件的所有需求模板](https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_template_list.html)

### 用户

- [x] [获取项目角色ID对照关系](https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_roles.html)

## 📄 License

[MIT](LICENSE)