package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath" // 👈 新增：用來處理路徑

	"github.com/Novie131/Nelo/internal/models"
	"github.com/Novie131/Nelo/internal/scanner"
)

// PushContext 掃描當前目錄並發送到指定的 Hub URL
func PushContext(serverURL string) error { // 👈 移除 projectName 參數
	// 1. 自動獲取當前工作目錄與資料夾名稱
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("獲取工作目錄失敗: %w", err)
	}
	projectName := filepath.Base(cwd) // 取得最後一個資料夾名稱

	// 2. 使用 scanner 掃描檔案
	excludes := []string{".git", "vendor", "bin", "go.sum", "project_context.md"}
	files, err := scanner.ScanProject(".", excludes)
	if err != nil {
		return fmt.Errorf("掃描失敗: %w", err)
	}

	// 3. 獲取當前機器名稱
	hostname, _ := os.Hostname()

	// 4. 封裝成 models 定義的結構
	payload := models.ProjectPayload{
		ProjectName: projectName,
		MachineName: hostname,
		Files:       files,
	}

	// 5. 將資料轉為 JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON 編碼失敗: %w", err)
	}

	// 6. 發送 POST 請求到 Hub
	fmt.Printf("📂 偵測到專案 [%s]，正在推送 %d 個檔案...\n", projectName, len(files))
	resp, err := http.Post(serverURL+"/api/context", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("連線至 Hub 失敗: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Hub 回傳錯誤代碼: %d", resp.StatusCode)
	}

	fmt.Printf("🚀 成功將專案推送至 Hub [%s]！\n", serverURL)
	return nil
}
