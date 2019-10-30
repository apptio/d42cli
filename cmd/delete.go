package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an entry from Device42",
	Args:  cobra.MinimumNArgs(2),
	Long: `The delete command is used to delete a record stored in Device42.

Examples include deleting a device record:
d42cli delete device servername23
`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
