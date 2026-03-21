#!/bin/bash

# 1. 定義要搜尋的關鍵字 (與 build.sh 產出的檔名邏輯一致)
# 這裡搜尋包含 "nelo-" 且帶有 "server" 參數的程序
PID=$(ps -ef | grep "nelo-" | grep "server" | grep -v grep | awk '{print $2}')

if [ -z "$PID" ]; then
    echo "⚠️ 找不到正在運行的 Nelo 服務"
else
    echo "🛑 正在停止 Nelo Hub (PID: $PID)..."
    kill $PID
    
    # 等待程序完全結束
    sleep 1
    if ps -p $PID > /dev/null; then
        echo "⚠️ 程序未響應，強制關閉..."
        kill -9 $PID
    fi
    echo "✅ 服務已停止"
fi