package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var minifyCmd = &cobra.Command{
	Use:   "minify [file]",
	Short: "Minifies a JSON file",
	Long:  `Removes all whitespace and newlines from a JSON file to save space.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		minifyFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(minifyCmd)
}

func minifyFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var buffer bytes.Buffer
	err = json.Compact(&buffer, data)
	if err != nil {
		fmt.Printf("Error minifying JSON (is it valid?): %v\n", err)
		return
	}

	fmt.Println(buffer.String())
}
