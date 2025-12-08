package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var rawOutput bool

var rootCmd = &cobra.Command{
	Use:   "bra",
	Short: "The CLI for BrasilAPI",
	Long: `
Interact with BrasilAPI services directly from your terminal.
Check out more at: https://github.com/BrasilAPI/BrasilAPI`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().BoolVar(&rawOutput, "raw", false, "Output raw JSON response from the API")

	rootCmd.AddCommand(banksCmd)
	rootCmd.AddCommand(cepCmd)
}

func printRaw(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	fmt.Println(string(data))
	return nil
}
