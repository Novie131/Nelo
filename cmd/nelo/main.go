package main

import (
	"fmt"
	"log"

	"github.com/Novie131/Nelo/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("無法加載設定檔: %v", err)
	}

	fmt.Printf("%s 正在啟動，預計監聽埠: %s\n", cfg.App.Name, cfg.App.Port)
	fmt.Printf("已偵測到 Gemini 配置，準備使用模型: %s\n", cfg.Providers.Gemini.Model)

}
