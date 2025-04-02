package mcp

import (
	"log"
	"net/http"

	"github.com/go-tapd/tapd"
	"github.com/go-tapd/tapd/mcp/internal/tools"
	"github.com/go-tapd/tapd/mcp/internal/tools/story/template_list"
	"github.com/go-tapd/tapd/mcp/internal/tools/user/roles"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	workspaceID int
	mcpServer   *server.MCPServer
	tapdClient  *tapd.Client
}

var _ http.Handler = (*Server)(nil)

func NewServer(workspaceID int, client *tapd.Client, opts ...Option) (*Server, error) {
	o, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		workspaceID: workspaceID,
		tapdClient:  client,
		mcpServer:   server.NewMCPServer(o.name, tapd.Version()),
	}

	srv.registerTools()

	return srv, nil
}

func (s *Server) registerTools() {
	tools.RegisterTools(s.mcpServer,
		// greetings.NewTool(),
		roles.NewTool(s.workspaceID, s.tapdClient),
		template_list.NewTool(s.workspaceID, s.tapdClient),
	)
}

func (s *Server) ServerStdio() error {
	log.Println("Tapd MCP server is running")
	return server.ServeStdio(s.mcpServer)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Tapd MCP server is running")
	server.NewSSEServer(s.mcpServer).ServeHTTP(w, r)
}
