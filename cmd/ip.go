package cmd

import (
	"fmt"

	"github.com/apptio/d42cli/httphelper"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Retrieve an IP record from Device42",
	Args:  cobra.MinimumNArgs(1),
	Long: `This sub-command will specifically retrieve an IP address entry.
	
Example:
./d42cli get ip 192.168.1.2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(searchIP(args[0]))
	},
}

//func DoRequest(uri string, target string) string {
func searchIP(target string) string {
	targetResponse := httphelper.DoRequest("GET", "search/?query=ip&string=", target)

	return targetResponse
}

func init() {
	getCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
