package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var version string

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID (v1, v4, v6, v7)",
	Long: `Generates a UUID. 
Default is v7 (Time-ordered).
Options:
  v1: Time + MAC Address (Classic)
  v4: Random (Standard)
  v6: Reordered Time (Legacy DB friendly)
  v7: Time-ordered (Modern standard)`,
	Run: func(cmd *cobra.Command, args []string) {
		var id uuid.UUID
		var err error

		switch version {
		case "1":
			id, err = uuid.NewUUID()
		case "4":
			id, err = uuid.NewRandom()
		case "6":
			id, err = uuid.NewV6()
		case "7":
			id, err = uuid.NewV7()
		default:
			fmt.Printf("Error: Unknown version '%s'. Use 1, 4, 6, or 7.\n", version)
			os.Exit(1)
		}

		if err != nil {
			fmt.Printf("Error generating UUID: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(id.String())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	uuidCmd.Flags().StringVarP(&version, "version", "v", "7", "UUID version to generate (1, 4, 6, 7)")
}
