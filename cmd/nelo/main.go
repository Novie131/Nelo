package main

import (
	"fmt"
	"os"

	"github.com/Novie131/Nelo/internal/client"
	"github.com/Novie131/Nelo/internal/server"

	// 保留你原本的 scanner 和 formatter import
	"github.com/Novie131/Nelo/internal/formatter"
	"github.com/Novie131/Nelo/internal/scanner"
)

func main() {
	// 如果沒有輸入參數，提示用法
	if len(os.Args) < 2 {
		fmt.Println("用法: nelo [server|bundle]")
		return
	}

	command := os.Args[1]

	switch command {
	case "server":
		fmt.Println("🚀 啟動 Nelo Hub 伺服器 (Port: 8080)...")
		server.Start("8080")

	case "push":
		// 預設發送到 localhost，如果有帶參數則使用參數做為 URL
		serverURL := "http://localhost:8080"
		if len(os.Args) > 2 {
			serverURL = os.Args[2]
		}

		err := client.PushContext(serverURL) // 👈 不用再傳 projectName 了
		if err != nil {
			fmt.Printf("❌ 推送失敗: %v\n", err)
		}

	case "bundle":
		// 這裡是你原本寫好的本地打包邏輯
		fmt.Println("🔍 Nelo 正在打包本地專案上下文...")
		excludes := []string{".git", "vendor", "bin", "go.sum"}
		files, err := scanner.ScanProject(".", excludes)
		if err != nil {
			fmt.Printf("❌ 掃描失敗: %v\n", err)
			return
		}
		output := formatter.ToMarkdown(files)
		err = os.WriteFile("project_context.md", []byte(output), 0644)
		if err != nil {
			fmt.Printf("❌ 寫入失敗: %v\n", err)
			return
		}
		fmt.Println("✅ 打包完成！產出 project_context.md")

	default:
		fmt.Printf("❌ 未知指令: %s\n", command)
	}
}
