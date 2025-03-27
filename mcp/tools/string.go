package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/mark3labs/mcp-go/mcp"
)

var _ Tool = &StringTool{}

type StringTool struct{}

func NewStringTool() *StringTool {
	return &StringTool{}
}

func (h *StringTool) Tool() mcp.Tool {
	return mcp.NewTool("get_string_length",
		mcp.WithDescription("Get the length of the string"),
		mcp.WithString("string",
			mcp.Required(),
			mcp.Description("The string to get the length of"),
		),
	)
}

type StringResponse struct {
	Length int `json:"length"`
}

func (h *StringTool) Handle(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	str, ok := request.Params.Arguments["string"].(string)
	if !ok {
		slog.Error("invalid string",
			slog.String("string", fmt.Sprintf("%v", request.Params.Arguments["string"])),
		)
		return nil, fmt.Errorf("invalid string")
	}

	res := StringResponse{
		Length: len(str),
	}

	response, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		slog.Error("marshal response",
			slog.String("error", err.Error()),
		)
		return nil, fmt.Errorf("marshal response: %w", err)
	}

	return mcp.NewToolResultText(string(response)), nil
}
