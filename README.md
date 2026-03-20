## 專案進度與發展路線圖 | Development Progress and Roadmap

### 目前版本 | Current Version
**v1.0.0 (Stable MVP)**

---

### 已完成功能 (v1) | Completed Features (v1)

* **跨平台支援 | Cross-Platform Support**
    * 提供適用於 Linux, Windows 以及 Raspberry Pi 4 的編譯版本。
    * Binaries available for Linux, Windows, and Raspberry Pi 4.

* **自動專案掃描 | Automatic Project Scanning**
    * 自動識別當前工作目錄名稱並根據副檔名過濾原始碼內容。
    * Automatic identification of project names and file filtering based on extensions.

* **中央準備區 | Centralized Staging Area**
    * 提供常駐的儀表板介面，在輸出至 AI 前統一管理多個專案上下文。
    * A persistent dashboard interface to manage multiple project contexts before LLM export.

* **靈活輸出選項 | Flexible Export Options**
    * **下載實體檔案 (Download)**: 產生 Markdown 檔案，可直接作為 Gemini 或 GPT 的附件上傳。
    * **合併複製 (Clipboard)**: 支援勾選多個專案內容並合併複製為純文字，便於直接貼上對話框。
    * **Download**: Generate Markdown files for direct upload as attachments.
    * **Merge and Copy**: Select and combine multiple project contexts into a single clipboard string.

* **響應式儀表板 | Responsive Dashboard**
    * 使用 Gin Gonic 搭配 Tailwind CSS 打造的直觀操作介面。
    * Intuitive management interface built with Gin Gonic and Tailwind CSS.

---

### 未來發展計畫 (v2) | Future Roadmap (v2)

* **持久化儲存層 | Persistence Layer**
    * 導入 SQLite 資料庫，確保伺服器重啟後暫存區內容不遺失。
    * Integrate SQLite to preserve staged data across server restarts.

* **Token 估算功能 | Token Estimation**
    * 即時計算選取內容的 Token 數量，避免超過大型語言模型限制。
    * Real-time token count for selected contexts to prevent context window overflow.

* **自定義排除規則 | Custom Exclusions**
    * 支援使用設定檔定義各個專案專屬的掃描忽略清單。
    * Support for project-specific scan rules via configuration files.

* **介面優化 | UI Enhancements**
    * 新增深色與淺色模式切換功能，優化視覺體驗。
    * Implement Dark/Light mode switching for better user experience.