package provider

import "context"

// Provider 是所有 AI 模型必須遵循的標準介面
type Provider interface {
	Name() string
	// GetResponse 接收 Prompt 並回傳 AI 的回答
	GetResponse(ctx context.Context, prompt string) (string, error)
}
