# Project Context: Nelo (v1.0.0 Snapshot)

## Directory Structure
```text
README.md
docs/snapshots/Nelo_2026-03-22.md
cmd/nelo/main.go
internal/client/sender.go
internal/formatter/markdown.go
internal/models/project.go
internal/scanner/scanner.go
internal/server/router.go
scripts/build.sh
scripts/start.sh
scripts/stop.sh
static/js/main.js
templates/index.html
```

## 關鍵程式碼更新 (v1.0.0)

### File: internal/server/router.go (新增狀態 API)
```go
var LatestProject *models.ProjectPayload
var LastPushTime int64 // 記錄最後推送時間戳

func Start(port string) {
    r := gin.Default()
    // ... 
    r.POST("/api/context", func(c *gin.Context) {
        // ... 解析資料
        LatestProject = &payload
        LastPushTime = time.Now().UnixNano() // 更新時間戳
        c.JSON(http.StatusOK, gin.H{"message": "✅ 成功接收專案"})
    })

    r.GET("/api/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"last_push_time": LastPushTime})
    })
    // ...
}
```

### File: static/js/main.js (新增自動刷新邏輯)
```javascript
let currentPushTime = null;

// 自動檢查更新
function checkForUpdates() {
    fetch('/api/status')
        .then(response => response.json())
        .then(data => {
            if (currentPushTime === null) {
                currentPushTime = data.last_push_time;
            } else if (data.last_push_time !== currentPushTime) {
                window.location.reload(); // 偵測到變動則刷新
            }
        });
}

// 持久化準備區資料
window.addEventListener('beforeunload', () => {
    localStorage.setItem('nelo_staged_files', JSON.stringify(stagedFiles));
});

document.addEventListener('DOMContentLoaded', () => {
    const saved = localStorage.getItem('nelo_staged_files');
    if (saved) stagedFiles = JSON.parse(saved);
    renderStagingArea();
    setInterval(checkForUpdates, 1500); // 啟動輪詢
});
```

---
*此快照包含 v1.0.0 的核心邏輯更新，其餘 Scanner 與 Formatter 邏輯沿用 v1.0.0。*