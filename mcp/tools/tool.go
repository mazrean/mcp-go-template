package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type Tool interface {
	Tool() mcp.Tool
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
