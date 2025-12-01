package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var csvCmd = &cobra.Command{
	Use:   "csv [file]",
	Short: "Convert CSV to JSON (or JSON to CSV)",
	Long: `Automatically converts between CSV and JSON based on the file extension.
  - Input .csv  -> Outputs JSON
  - Input .json -> Outputs CSV (Must be an array of objects)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		convertFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)
}

func convertFile(filename string) {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".csv":
		csvToJson(filename)
	case ".json":
		jsonToCsv(filename)
	default:
		fmt.Println("Error: File must be .csv or .json")
	}
}

func csvToJson(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(records) < 1 {
		fmt.Println("CSV is empty")
		return
	}

	headers := records[0]
	var result []map[string]string

	for _, row := range records[1:] {
		record := make(map[string]string)
		for i, value := range row {
			if i < len(headers) {
				record[headers[i]] = value
			}
		}
		result = append(result, record)
	}

	jsonData, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonData))
}

func jsonToCsv(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Println("Error: JSON must be an array of objects (e.g. [{},{}])")
		return
	}

	if len(data) == 0 {
		fmt.Println("JSON array is empty")
		return
	}

	var headers []string
	for k := range data[0] {
		headers = append(headers, k)
	}

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write(headers)

	for _, obj := range data {
		var row []string
		for _, header := range headers {
			val := obj[header]
			row = append(row, fmt.Sprintf("%v", val))
		}
		writer.Write(row)
	}
}
