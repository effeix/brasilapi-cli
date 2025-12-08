package cli

import (
	"fmt"
	"regexp"

	"github.com/effeix/brasilapi-cli/internal/api"
	"github.com/spf13/cobra"
)

var cepCmd = &cobra.Command{
	Use:   "cep [cep]",
	Short: "Get address information by Brazilian postal code (CEP)",
	Long: `
Fetches address information for a given Brazilian postal code (CEP).

The CEP must be exactly 8 digits. You can provide it with or without
the hyphen.

Examples:
  > Get address information for CEP 01001000
  bra cep 01001000

  > Get address information for CEP 01001-000
  bra cep 01001-000
`,
	Args: cobra.ExactArgs(1),
	RunE: runCEP,
}

func runCEP(cmd *cobra.Command, args []string) error {
	rawOutput, _ := cmd.Root().PersistentFlags().GetBool("raw")

	cep := sanitizeCEP(args[0])
	if !isValidCEP(cep) {
		return fmt.Errorf("invalid CEP: must be exactly 8 digits")
	}

	client := api.NewClient()
	result, err := client.GetCEP(cep)
	if err != nil {
		if _, ok := err.(*api.BrasilAPIError); ok && rawOutput {
			return printRaw(err)
		}
		return fmt.Errorf("failed to fetch CEP: %w", err)
	}

	if rawOutput {
		return printRaw(result)
	}

	displayCEP(result)
	return nil
}

var regex = regexp.MustCompile(`\D`)

func sanitizeCEP(cep string) string {
	return regex.ReplaceAllString(cep, "")
}

func isValidCEP(cep string) bool {
	return len(cep) == 8
}

func displayCEP(cep *api.CEP) {
	fmt.Printf("CEP:          %s\n", cep.CEP)
	fmt.Printf("State:        %s\n", cep.State)
	fmt.Printf("City:         %s\n", cep.City)
	fmt.Printf("Neighborhood: %s\n", cep.Neighborhood)
	fmt.Printf("Street:       %s\n", cep.Street)
	fmt.Printf("Service:      %s\n", cep.Service)
}
