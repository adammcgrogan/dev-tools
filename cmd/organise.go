package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Sorts files in the current directory by extension",
	Long:  `Scans the current directory and moves files into subfolders (Images, Docs, etc.) based on their file extension.`,
	Run: func(cmd *cobra.Command, args []string) {
		organizeFiles()
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}

func organizeFiles() {
	// 1. read current directory
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Scanning files...")

	for _, file := range files {
		// skip directories and hidden files
		if file.IsDir() || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		// skip the tool itself
		if file.Name() == "main.go" || file.Name() == "dev-tools" || file.Name() == "go.mod" || file.Name() == "go.sum" {
			continue
		}

		// determine Category
		ext := strings.ToLower(filepath.Ext(file.Name()))
		folder := getCategory(ext)

		if folder == "" {
			fmt.Printf("Skipping: %s (unknown type)\n", file.Name())
			continue
		}

		// 3. create folder if it doesn't exist
		if err := os.MkdirAll(folder, 0755); err != nil {
			fmt.Printf("Error creating folder %s: %v\n", folder, err)
			continue
		}

		// 4. move the file
		oldPath := file.Name()
		newPath := filepath.Join(folder, file.Name())

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Failed to move %s: %v\n", file.Name(), err)
		} else {
			fmt.Printf("Moved: %s -> %s/\n", file.Name(), folder)
		}
	}
}

func getCategory(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".svg", ".webp":
		return "Images"
	case ".pdf", ".doc", ".docx", ".txt", ".md", ".xlsx", ".csv":
		return "Documents"
	case ".mp3", ".wav", ".mp4", ".mov", ".avi":
		return "Media"
	case ".zip", ".tar", ".gz", ".rar":
		return "Archives"
	case ".exe", ".dmg", ".pkg", ".deb":
		return "Installers"
	default:
		return "" // Unknown types stay put
	}
}
