package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var decode bool

var base64Cmd = &cobra.Command{
	Use:   "base64 [text]",
	Short: "Encode or decode Base64 strings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		if decode {
			data, err := base64.StdEncoding.DecodeString(input)
			if err != nil {
				fmt.Println("Error: Invalid Base64 string")
				return
			}
			fmt.Println(string(data))
		} else {
			encoded := base64.StdEncoding.EncodeToString([]byte(input))
			fmt.Println(encoded)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().BoolVarP(&decode, "decode", "d", false, "Decode input instead of encoding")
}
