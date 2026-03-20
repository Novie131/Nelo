package formatter

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Novie131/Nelo/internal/scanner"
)

// ToMarkdown 將檔案列表轉換為一個大 Markdown 文本
func ToMarkdown(files []scanner.FileInfo) string {
	var builder strings.Builder

	builder.WriteString("# Project Context: Nelo\n\n")

	// 1. 先產出目錄結構 (Tree)
	builder.WriteString("## Directory Structure\n```text\n")
	for _, f := range files {
		builder.WriteString(f.Path + "\n")
	}
	builder.WriteString("```\n\n")

	// 2. 產出每個檔案的代碼塊
	builder.WriteString("## File Contents\n\n")
	for _, f := range files {
		builder.WriteString(fmt.Sprintf("### File: %s\n", f.Path))
		// 根據副檔名決定代碼塊語法
		ext := strings.TrimPrefix(filepath.Ext(f.Path), ".")
		builder.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", ext, f.Content))
	}

	return builder.String()
}
