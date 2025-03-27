// Package tools provides tools for the Model Context Protocol (MCP) server.
//
// The tools can be added to the server with the NewServer function.
package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// Tool represents a tool that can be added to the Model Context Protocol (MCP) server.
type Tool interface {
	Tool() mcp.Tool
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
