package server

import (
	"net/http"

	"github.com/Novie131/Nelo/internal/formatter"
	"github.com/Novie131/Nelo/internal/models"
	"github.com/gin-gonic/gin"
)

// LatestProject 暫存最新收到的專案
var LatestProject *models.ProjectPayload

// Start 啟動 Nelo Hub
func Start(port string) {
	r := gin.Default()

	// 載入 HTML 模板
	r.LoadHTMLGlob("templates/*")

	// 👈 關鍵修正：開放 static 資料夾給瀏覽器讀取 JS 與 CSS
	r.Static("/static", "./static")

	// 📡 API 1: 接收 Client 推送過來的專案資料
	r.POST("/api/context", func(c *gin.Context) {
		var payload models.ProjectPayload

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "資料解析失敗: " + err.Error()})
			return
		}

		LatestProject = &payload

		c.JSON(http.StatusOK, gin.H{
			"message":    "✅ 成功接收專案: " + payload.ProjectName,
			"file_count": len(payload.Files),
			"from":       payload.MachineName,
		})
	})

	// 📊 API 2: 渲染 Dashboard 網頁介面
	r.GET("/dashboard", func(c *gin.Context) {
		if LatestProject == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Project": nil,
			})
			return
		}

		markdownContent := formatter.ToMarkdown(LatestProject.Files)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Project":  LatestProject,
			"Markdown": markdownContent,
		})
	})

	// 🧹 API 3: 清空 Hub 資料
	r.DELETE("/api/context", func(c *gin.Context) {
		LatestProject = nil
		c.JSON(http.StatusOK, gin.H{"message": "🗑️ Hub 已經清空"})
	})

	r.Run(":" + port)
}
