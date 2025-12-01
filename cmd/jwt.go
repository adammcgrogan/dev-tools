package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt [token]",
	Short: "Decodes a JWT token",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		decodeToken(args[0])
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

func decodeToken(token string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Error: Invalid JWT format (needs 3 parts)")
		return
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}

	// unmarshal into a map so we can add fields to it
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// check expiry and inject into the JSON map
	if exp, ok := claims["exp"].(float64); ok {
		expiryTime := time.Unix(int64(exp), 0)

		// human readable date to the JSON
		claims["expires_human"] = expiryTime.Format(time.RFC3339)

		// add boolean validity to the JSON
		if time.Now().After(expiryTime) {
			claims["valid"] = false
		} else {
			claims["valid"] = true
		}
	}

	output, _ := json.MarshalIndent(claims, "", "  ")
	fmt.Println(string(output))
}
