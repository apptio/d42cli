package cmd

import (
	"fmt"
	"os"

	"github.com/apptio/d42cli/httphelper"
	"github.com/spf13/cobra"
)

var allDevices bool

// deviceCmd represents the device command
var getDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Retrieve a device record from Device42",
	Long: `This sub-command will retrieve an entry of type 'device'.

Examples:
d42cli get device servername23
d42cli get device --all #WARNING: this is a large API call to d42
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if allDevices {
			// All devices flag detected
			fmt.Println(getAllDevice())
		} else if len(args) > 0 {
			// We are after a single device
			fmt.Println(getDevice(args[0]))
		} else {
			// No flag or device target passed - output help
			cmd.Help()
			os.Exit(0)
		}
	},
}

func getDevice(target string) string {
	targetResponse := httphelper.DoRequest("GET", "devices/name/", target+"/")

	return targetResponse
}

func getAllDevice() string {
	targetResponse := httphelper.DoRequest("GET", "devices", "/")

	return targetResponse
}

func init() {
	getCmd.AddCommand(getDeviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getDeviceCmd.Flags().BoolVarP(&allDevices, "all", "a", false, "Request all devices with brief information")
}
