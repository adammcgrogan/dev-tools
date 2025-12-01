package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree [dir]",
	Short: "Visualizes directory structure",
	Long:  `Recursively prints the file structure of the current (or specified) directory in a tree format.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := "."
		if len(args) > 0 {
			root = args[0]
		}
		fmt.Println(root)
		printTree(root, "")
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}

func printTree(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return // Skip unreadable folders
	}

	var filtered []os.DirEntry
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" {
			continue
		}
		filtered = append(filtered, e)
	}

	for i, entry := range filtered {
		isLast := i == len(filtered)-1

		connector := "├── "
		childPrefix := "│   "
		if isLast {
			connector = "└── "
			childPrefix = "    "
		}

		fmt.Println(prefix + connector + entry.Name())

		if entry.IsDir() {
			newPath := filepath.Join(path, entry.Name())
			printTree(newPath, prefix+childPrefix)
		}
	}
}
