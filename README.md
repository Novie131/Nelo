# Nelo - Project Context Bundler & Hub

Nelo 是一個專為大型語言模型（LLM）開發者設計的上下文管理工具。它能自動掃描專案原始碼，並透過中央 Hub 進行管理、合併與輸出，讓你更輕鬆地將程式碼餵給 AI [cite: Nelo_1-21-41.md]。

---

## 專案進度與發展路線圖 | Development Progress and Roadmap

### 目前版本 | Current Version
**v1.0.0 (Stable MVP)** [cite: Nelo_1-21-41.md]

---

### 已完成功能 (v1) | Completed Features (v1)

* **背景常駐運行 | Background Operation**
    * 支援透過 `nohup` 腳本實現服務不中斷運行，並自動導向日誌紀錄 [cite: Nelo_1-21-41.md]。
* **跨平台支援 | Cross-Platform Support**
    * 提供適用於 Linux, Windows 以及 Raspberry Pi 4 的編譯版本 [cite: Nelo_1-21-41.md]。
* **自動專案掃描 | Automatic Project Scanning**
    * 自動識別當前工作目錄名稱並根據副檔名過濾原始碼內容 [cite: Nelo_1-21-41.md]。
* **中央準備區 | Centralized Staging Area**
    * 提供常駐的儀表板介面，在輸出至 AI 前統一管理多個專案上下文 [cite: Nelo_1-21-41.md]。
* **靈活輸出選項 | Flexible Export Options**
    * **下載實體檔案 (Download)**: 產生 Markdown 檔案，可直接作為 Gemini 或 GPT 的附件上傳 [cite: Nelo_1-21-41.md]。
    * **合併複製 (Clipboard)**: 支援勾選多個專案內容並合併複製為純文字，便於直接貼上對話框 [cite: Nelo_1-21-41.md]。
* **響應式儀表板 | Responsive Dashboard**
    * 使用 Gin Gonic 搭配 Tailwind CSS 打造的直觀操作介面 [cite: Nelo_1-21-41.md]。

---

### 未來發展計畫 (v2) | Future Roadmap (v2)

* **持久化儲存層 | Persistence Layer**
    * 導入 SQLite 資料庫，確保伺服器重啟後暫存區內容不遺失 [cite: Nelo_1-21-41.md]。
* **Token 估算功能 | Token Estimation**
    * 即時計算選取內容的 Token 數量，避免超過大型語言模型限制 [cite: Nelo_1-21-41.md]。
* **自定義排除規則 | Custom Exclusions**
    * 支援使用設定檔定義各個專案專屬的掃描忽略清單 [cite: Nelo_1-21-41.md]。
* **介面優化 | UI Enhancements**
    * 新增深色與淺色模式切換功能，優化視覺體驗 [cite: Nelo_1-21-41.md]。

---

## 安裝與執行說明 | Installation & Execution

### 1. 初始化環境
在使用腳本前，請先賦予執行權限並安裝全域指令：

```bash
# 賦予所有腳本執行權限 [cite: Nelo_1-21-41.md]
chmod +x scripts/*.sh

# 安裝 nelo 指令到系統路徑，方便在任何專案目錄執行 push
go install ./cmd/nelo
```

### 2. 管理 Nelo Hub (伺服器端)
Hub 負責接收來自各個專案的資料並提供網頁介面。

* **編譯專案** [cite: Nelo_1-21-41.md]
  ```bash
  ./scripts/build.sh
  ```
* **啟動背景服務** [cite: Nelo_1-21-41.md]
  ```bash
  ./scripts/start.sh
  ```
* **查看運行日誌** [cite: Nelo_1-21-41.md]
  ```bash
  tail -f nelo.log
  ```
* **停止服務** [cite: Nelo_1-21-41.md]
  ```bash
  ./scripts/stop.sh
  ```
* **訪問儀表板**：瀏覽器開啟 `http://localhost:8080/dashboard` [cite: Nelo_1-21-41.md]。

### 3. 推送專案 (用戶端)
當 Hub 啟動後，你可以在任何專案目錄快速推送內容：

```bash
# 在目標專案資料夾執行 (需先執行 go install)
nelo push

# 或是手動指定 Hub 位址
nelo push http://localhost:8080
```

---

## 專案結構 | Directory Structure
```text
cmd/nelo/          # 程式進入點
internal/          # 內部邏輯 (Scanner, Formatter, Server)
scripts/           # 自動化管理腳本
static/ & templates/ # 網頁前端資源