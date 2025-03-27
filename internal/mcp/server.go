// Package mcp provides a way to create a Model Context Protocol (MCP) server.
//
// The server can be started with the Start method.
package mcp

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/mazrean/mcp-tarmaq/mcp/tools"
)

// Server represents a Model Context Protocol (MCP) server.
type Server struct {
	server *server.MCPServer
}

// NewServer creates a new Model Context Protocol (MCP) server.
func NewServer(
	version string,
	tools ...tools.Tool,
) *Server {
	s := server.NewMCPServer(
		"mcp-go-template", // Name
		version,           // Version
		server.WithLogging(),
	)

	for _, tool := range tools {
		s.AddTool(tool.Tool(), tool.Handle)
	}

	return &Server{
		server: s,
	}
}

// Start starts the Model Context Protocol (MCP) server.
func (s *Server) Start() error {
	return server.ServeStdio(s.server)
}
