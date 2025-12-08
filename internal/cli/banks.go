package cli

import (
	"fmt"
	"regexp"

	"github.com/effeix/brasilapi-cli/internal/api"
	"github.com/spf13/cobra"
)

var banksCmd = &cobra.Command{
	Use:   "banks [code]",
	Short: "List Brazilian banks or get details of a specific bank",
	Long: `
Fetches the list of all Brazilian banks or details of a specific bank by its COMPE code.

When called without arguments, lists all available banks.

Examples:
  > List all banks
  bra banks

  > Get details for Banco do Brasil (code 1)
  bra banks 1

  > Get details for Itaú (code 341)
  bra banks 341
`,
	Args: cobra.MaximumNArgs(1),
	RunE: runBanks,
}

var regexOnlyDigits = regexp.MustCompile(`^\d+$`)

func runBanks(cmd *cobra.Command, args []string) error {
	rawOutput, _ := cmd.Root().PersistentFlags().GetBool("raw")

	client := api.NewClient()

	if len(args) == 0 {
		return listAllBanks(client, rawOutput)
	}

	code := args[0]
	if !regexOnlyDigits.MatchString(code) {
		return fmt.Errorf("bank code must contain only numeric digits")
	}
	if len(code) > 3 {
		return fmt.Errorf("bank code must have up to 3 digits")
	}

	return getBankByCode(client, code, rawOutput)
}

func listAllBanks(client *api.Client, rawOutput bool) error {
	banks, err := client.GetBanks()
	if err != nil {
		if _, ok := err.(*api.BrasilAPIError); ok && rawOutput {
			return printRaw(err)
		}
		return fmt.Errorf("failed to fetch banks: %w", err)
	}

	sortBanks(banks)

	if rawOutput {
		return printRaw(banks)
	}

	displayBanks(banks)
	return nil
}

func getBankByCode(client *api.Client, code string, rawOutput bool) error {
	bank, err := client.GetBankByCode(code)
	if err != nil {
		if _, ok := err.(*api.BrasilAPIError); ok && rawOutput {
			return printRaw(err)
		}
		return fmt.Errorf("failed to fetch bank: %w", err)
	}

	if rawOutput {
		return printRaw(bank)
	}

	displayBank(bank)
	return nil
}

func displayBank(bank *api.Bank) {
	if bank == nil {
		fmt.Println("Cannot display bank: no data available")
		return
	}

	fmt.Printf("Code:      %03d\n", bank.Code)
	fmt.Printf("ISPB:      %s\n", bank.ISPB)
	fmt.Printf("Name:      %s\n", bank.Name)
	fmt.Printf("Full Name: %s\n", bank.FullName)
}

func displayBanks(banks []*api.Bank) {
	fmt.Printf("%-6s %-11s %s\n", "CODE", "ISPB", "FULL NAME")
	fmt.Println("--------------------------------------------------------------")

	for _, bank := range banks {
		fmt.Printf("%-6.3d %-11s %s\n", bank.Code, bank.ISPB, bank.FullName)
	}

	fmt.Printf("\nTotal: %d banks\n", len(banks))
}

func sortBanks(banks []*api.Bank) {
	for i := 0; i < len(banks)-1; i++ {
		for j := i + 1; j < len(banks); j++ {
			if banks[i].Code > banks[j].Code {
				banks[i], banks[j] = banks[j], banks[i]
			}
		}
	}
}
