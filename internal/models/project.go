package models

import "github.com/Novie131/Nelo/internal/scanner"

// ProjectPayload 是發送到 Server 的資料格式
type ProjectPayload struct {
	ProjectName string             `json:"project_name"`
	MachineName string             `json:"machine_name"` // 辨識是 Pi 4 還是 Windows
	Files       []scanner.FileInfo `json:"files"`
}
