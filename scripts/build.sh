#!/bin/bash
# 執行指令前請先 chmod +x scripts/build.sh

# 遇到錯誤立刻停止執行
set -e

# 確保輸出目錄存在
mkdir -p bin

# 利用 go env 自動獲取當前機器的作業系統與架構
HOST_OS=$(go env GOOS)
HOST_ARCH=$(go env GOARCH)

# 處理 Windows 特殊的副檔名需求
EXT=""
if [ "$HOST_OS" == "windows" ]; then
    EXT=".exe"
fi

# 動態組合檔名 (例如: nelo-linux-amd64 或 nelo-windows-amd64.exe)
OUTPUT_NAME="nelo-${HOST_OS}-${HOST_ARCH}${EXT}"

echo "🔍 偵測到當前系統架構: ${HOST_OS} (${HOST_ARCH})"
echo "🔨 開始編譯專屬版本..."

# 執行編譯 (不特別指定 GOOS/GOARCH，Go 預設就會編譯當前系統的版本)
go build -o "bin/${OUTPUT_NAME}" ./cmd/nelo

echo "✅ 編譯完成！"
echo "📂 檔案位置: bin/${OUTPUT_NAME}"