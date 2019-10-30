package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an entry from Device42",
	Args:  cobra.MinimumNArgs(2),
	Long: `The get command is used to retrieve details of a record stored in Device42.

Examples include retrieving IP records, device records:
d42cli get device servername23
d42cli get ip 172.0.1.1
`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
