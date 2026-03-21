#!/bin/bash

# 1. 取得腳本所在的目錄，並定位到專案根目錄 (scripts 的上一層)
SCRIPT_DIR=$(cd "$(dirname "$0")"; pwd)
PROJECT_ROOT=$(cd "$SCRIPT_DIR/.."; pwd)
cd "$PROJECT_ROOT"

# 2. 搜尋編譯好的執行檔 (通常在 bin 目錄下)
# 會自動抓取第一個 nelo- 開頭的檔案
EXEC_FILE=$(ls bin/nelo-* 2>/dev/null | head -n 1)

if [ -z "$EXEC_FILE" ]; then
    echo "❌ 錯誤：在 bin/ 找不到執行檔。請先執行 ./scripts/build.sh"
    exit 1
fi

# 3. 檢查是否已經有在執行的服務，避免重複啟動
PID=$(ps -ef | grep "$EXEC_FILE server" | grep -v grep | awk '{print $2}')
if [ -n "$PID" ]; then
    echo "⚠️ Nelo Hub 已經在運行中 (PID: $PID)"
    exit 0
fi

# 4. 使用 nohup 在背景啟動
# 將日誌輸出到根目錄的 nelo.log
nohup ./"$EXEC_FILE" server > nelo.log 2>&1 &

# 5. 輸出結果
NEW_PID=$!
echo "🚀 Nelo Hub 已在背景啟動 (PID: $NEW_PID)"
echo "📊 日誌位置：$PROJECT_ROOT/nelo.log"