package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "ping [url1] [url2]...",
	Short: "Check the status of websites",
	Long:  `Checks the availability and response time of multiple websites concurrently.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runChecks(args)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func runChecks(urls []string) {
	fmt.Printf("Checking %d websites...\n\n", len(urls))

	var wg sync.WaitGroup
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)

		go func(site string) {
			defer wg.Done()
			checkUrl(site)
		}(url)
	}

	wg.Wait()
	fmt.Printf("\nDone in %v\n", time.Since(start))
}

func checkUrl(url string) {
	// Add https if missing
	target := url
	if len(url) < 4 || url[:4] != "http" {
		target = "https://" + url
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(target)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("[DOWN] %s (Error: %v)\n", target, err)
		return
	}
	defer resp.Body.Close()

	status := "UP"
	if resp.StatusCode != 200 {
		status = fmt.Sprintf("WARN (%d)", resp.StatusCode)
	}

	fmt.Printf("[%s] %s (%v)\n", status, target, duration)
}
