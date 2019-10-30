package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/apptio/d42cli/httphelper"
	"github.com/spf13/cobra"
)

var forceDelete bool

// Device42Device Struct to contain device42 device JSON
type Device42Device struct {
	LastUpdated time.Time `json:"last_updated"`
	Orientation int       `json:"orientation"`
	IPAddresses []struct {
		SubnetID int    `json:"subnet_id"`
		IP       string `json:"ip"`
		Label    string `json:"label"`
		Type     int    `json:"type"`
		Subnet   string `json:"subnet"`
	} `json:"ip_addresses"`
	SerialNo      string      `json:"serial_no"`
	HwDepth       interface{} `json:"hw_depth"`
	DeviceID      int         `json:"device_id"`
	ServiceLevel  string      `json:"service_level"`
	IsItBladeHost string      `json:"is_it_blade_host"`
	HwSize        interface{} `json:"hw_size"`
	ID            int         `json:"id"`
	CustomFields  []struct {
		Notes interface{} `json:"notes"`
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"custom_fields"`
	Aliases                 []interface{} `json:"aliases"`
	Category                string        `json:"category"`
	HddDetails              interface{}   `json:"hdd_details"`
	UUID                    string        `json:"uuid"`
	VirtualSubtype          string        `json:"virtual_subtype"`
	Cpuspeed                interface{}   `json:"cpuspeed"`
	HwModel                 interface{}   `json:"hw_model"`
	HddraidType             interface{}   `json:"hddraid_type"`
	RackID                  int           `json:"rack_id"`
	Hddcount                interface{}   `json:"hddcount"`
	Building                string        `json:"building"`
	Xpos                    int           `json:"xpos"`
	DeviceExternalLinks     []interface{} `json:"device_external_links"`
	StartAt                 float64       `json:"start_at"`
	Tags                    []interface{} `json:"tags"`
	InService               bool          `json:"in_service"`
	Hddsize                 interface{}   `json:"hddsize"`
	MacAddresses            []interface{} `json:"mac_addresses"`
	Hddraid                 interface{}   `json:"hddraid"`
	Cpucount                interface{}   `json:"cpucount"`
	Os                      interface{}   `json:"os"`
	VirtualHostName         string        `json:"virtual_host_name"`
	IsItVirtualHost         string        `json:"is_it_virtual_host"`
	IsItSwitch              string        `json:"is_it_switch"`
	Customer                interface{}   `json:"customer"`
	UcsManager              interface{}   `json:"ucs_manager"`
	Name                    string        `json:"name"`
	Room                    string        `json:"room"`
	Row                     string        `json:"row"`
	Type                    string        `json:"type"`
	Notes                   string        `json:"notes"`
	RAM                     interface{}   `json:"ram"`
	AssetNo                 string        `json:"asset_no"`
	Manufacturer            interface{}   `json:"manufacturer"`
	DevicePurchaseLineItems []interface{} `json:"device_purchase_line_items"`
	Cpucore                 interface{}   `json:"cpucore"`
	Where                   int           `json:"where"`
	Rack                    string        `json:"rack"`
}

// deviceCmd represents the device command
var deleteDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Delete a device record from Device42",
	Long: `This sub-command will delete an entry of type 'device'.

Examples:
d42cli delete device servername23
d42cli delete device servername23 -f
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if forceDelete {
			// All devices flag detected
			fmt.Println(deleteDevice(args[0], forceDelete))
		} else if len(args) > 0 {
			// Delete a single device
			fmt.Println(deleteDevice(args[0], false))
		} else {
			// No flag or device target passed - output help
			cmd.Help()
			os.Exit(0)
		}
	},
}

func deleteDevice(target string, force bool) string {
	// To delete a device, we must first find its ID
	targetResponse := httphelper.DoRequest("GET", "devices/name/", target+"/")

	if !strings.Contains(targetResponse, "last_updated") {
		// The call has failed, return the message
		return targetResponse
	}

	// Unmarshal the response and extract the device_id
	d42Device := Device42Device{}
	json.Unmarshal([]byte(targetResponse), &d42Device)
	deviceID := strconv.Itoa(d42Device.DeviceID)

	if !force {
		fmt.Println("Confirm delete of object with ID " + deviceID + ": y/n?")
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		if strings.Compare("y", answer) == 0 {
			deleteResponse := doDeleteDevice(deviceID)
			return deleteResponse
		}

		return "{\"msg\": \"Delete command aborted\"}"
	}

	deleteResponse := doDeleteDevice(deviceID)
	return deleteResponse
}

func doDeleteDevice(targetID string) string {
	deleteResponse := httphelper.DoRequest("DELETE", "devices/", targetID+"/")

	return deleteResponse
}

func init() {
	deleteCmd.AddCommand(deleteDeviceCmd)

	deleteDeviceCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Do not prompt for confirmation during delete.")
}
