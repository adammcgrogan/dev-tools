package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID v7 (Time-ordered)",
	Long:  `Generates a Version 7 UUID. Unlike v4, v7 is time-ordered, making it sortable and database-friendly.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := uuid.NewV7()
		if err != nil {
			fmt.Printf("Error generating UUID: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(id.String())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
