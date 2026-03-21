# Nelo - AI Context Hub

Nelo 是一個專為大型語言模型（LLM）開發者設計的上下文管理工具。它能自動掃描專案原始碼，並透過中央 Hub 進行管理、合併與輸出，讓你更輕鬆地將程式碼餵給 AI。

---

## 目前版本 | Current Version: **v1.1.0**

### 更新亮點 (v1.1.0)
* **自動刷新 Dashboard**: 實作前端輪詢機制，`push` 後網頁會自動更新，無需手動按 F5。
* **狀態持久化**: 使用 `localStorage` 技術，確保網頁重整後「準備區」的資料不會遺失。
* **結構優化**: 將詳細的開發快照移至 `docs/` 資料夾，保持根目錄整潔。

---

## 更新日誌與快照 | Changelog & Snapshots

為了讓 AI 能快速理解專案演進，我們將各版本的完整上下文備份於此：

| 版本 | 日期 | 重點內容 | 詳細快照 (Context) |
| :--- | :--- | :--- | :--- |
| **v1.1.0** | 2026-03-22 | 實作自動刷新、解決快取問題 | [詳細內容](./docs//Nelo_2026-03-22.md) |

---

## 安裝與執行 | Installation & Execution

### 1. 初始化環境
```bash
# 賦予腳本執行權限
chmod +x scripts/*.sh

# 安裝 nelo 指令
go install ./cmd/nelo
```

### 2. 啟動 Nelo Hub (伺服器端)
```bash
./scripts/build.sh
./scripts/start.sh
```
* **瀏覽器開啟**: `http://localhost:8080/dashboard`

### 3. 推送專案 (用戶端)
在目標專案資料夾執行：
```bash
nelo push
```

### 4. 懶人指令
```bash
bash ./scripts/stop.sh && bash ./scripts/build.sh && bash ./scripts/start.sh
```

---

## 專案結構 | Directory Structure
* `cmd/nelo/`: 程式進入點。
* `docs/snapshots/`: 存放專案演進的 Markdown 快照。
* `internal/`: 核心邏輯（Scanner, Server, Router）。
* `scripts/`: 自動化管理腳本。
* `static/ & templates/`: 網頁前端資源。