package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var epochCmd = &cobra.Command{
	Use:   "epoch [timestamp]",
	Short: "Convert Unix timestamp to Date (or get current)",
	Long:  `Converts a Unix timestamp (seconds) to a human-readable date. If no argument is provided, returns the current timestamp.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			now := time.Now()
			fmt.Printf("Current Timestamp: %d\n", now.Unix())
			fmt.Printf("Current Date:      %s\n", now.Format(time.RFC1123))
			return
		}

		input := args[0]
		i, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Error: Timestamp must be a valid integer")
			return
		}

		tm := time.Unix(i, 0)

		fmt.Printf("Input:  %d\n", i)
		fmt.Printf("UTC:    %s\n", tm.UTC().Format(time.RFC1123))
		fmt.Printf("Local:  %s\n", tm.Local().Format(time.RFC1123))
		fmt.Printf("ISO:    %s\n", tm.Format(time.RFC3339))
	},
}

func init() {
	rootCmd.AddCommand(epochCmd)
}
