package cmd

import (
	"os"

	"github.com/jeffersonayub/goexpert-stress-test/internal/infraestrutura"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goexpert-stress-test",
	Short: "Desafio Stress Test do curso GoExpert",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")

		if err != nil || url == "" {
			println("Error: URL is required")
			os.Exit(1)
		}

		requests, err := cmd.Flags().GetInt("requests")
		if err != nil || requests <= 0 {
			println("Error: Invalid number of requests", requests)
			os.Exit(1)
		}

		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil || concurrency <= 0 {
			println("Error: Invalid number of concurrent requests")
			os.Exit(1)
		}

		if concurrency > requests {
			println("Error: Concurrency is greater than requests")
			os.Exit(1)
		}

		infraestrutura.StressTest(url, requests, concurrency)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to process")
	rootCmd.Flags().IntP("requests", "r", 0, "Number of requests to send")
	rootCmd.Flags().IntP("concurrency", "c", 0, "Number of concurrent requests")
}
