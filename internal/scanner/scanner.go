package scanner

import (
	"io" // 👈 新增 io 套件來判斷 io.EOF
	"os"
	"path/filepath"
	"strings"
)

// FileInfo 儲存檔案路徑與內容
type FileInfo struct {
	Path    string
	Content string
}

// ScanProject 掃描目錄並回傳符合條件的檔案列表
func ScanProject(root string, excludes []string) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 跳過隱藏資料夾（如 .git）或指定的排除路徑
		for _, ex := range excludes {
			if strings.Contains(path, ex) {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// 💡 新增邏輯：如果是資料夾，檢查是否為空
		if d.IsDir() {
			f, err := os.Open(path)
			if err == nil {
				// 嘗試讀取 1 個內容
				_, err = f.Readdirnames(1)
				f.Close()

				// 如果回傳 io.EOF，代表裡面完全沒有檔案或子目錄（空資料夾）
				if err == io.EOF {
					relPath, _ := filepath.Rel(root, path)
					files = append(files, FileInfo{
						Path:    relPath + "/",       // 結尾加上斜線，讓肉眼一看就知道是資料夾
						Content: "(Empty Directory)", // 塞入提示文字取代原本的程式碼內容
					})
				}
			}
			return nil // 資料夾處理完畢，繼續往下掃描
		}

		// 原本邏輯：只抓取我們感興趣的程式碼檔案
		if isTargetFile(path) {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil // 讀不到就跳過
			}
			relPath, _ := filepath.Rel(root, path)
			files = append(files, FileInfo{
				Path:    relPath,
				Content: string(content),
			})
		}
		return nil
	})

	return files, err
}

func isTargetFile(path string) bool {
	exts := []string{
		".go", ".yaml", ".yml", ".sh", ".md", ".json",
		".js", ".ts", ".css", ".html", ".htm",
		".py", ".c", ".cpp", ".h",
	}

	for _, ext := range exts {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
