# d42cli

d42cli is a simple tool to interact with the Device42 API.

## Configuration

Configuration is in the form of a yaml file in the users homedir - ".d42cli.yaml"

The contents of the users config file should be as follows:
```
{
    "Username": "USERNAME",
    "Password": "PASSWORD",
    "BaseURL": "https://device42.base.url/api/1.0/"
}
```

**Output**

```
This tool is designed to help interact with the Device42 CMDB appliance.
This binary interacts with your Device42 appliance via the 1.0 API.

Examples;
to retrieve a device entry called servername3:
d42cli get device servername3

Usage:
  d42cli [command]

Available Commands:
  delete      Delete an entry from Device42
  get         Retrieve an entry from Device42
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.d42cli.yaml)
  -h, --help            help for d42cli
  -t, --toggle          Help message for toggle

Use "d42cli [command] --help" for more information about a command.
```

Here is an example of getting a device by name and using `jq` to process the results:
```
./d42cli get device sample.device.name | jq .name,.room,.hw_model
"sample.device.name"
"DC1 Cage1"
"MANUF-MODEL-5s"
```

While available, its discouraged to get all devices. This call takes some time for d42 to reply depending on the size of your database:
```
./d42cli get device --all
```

You can also do the same for a few entry types;
```
./d42cli get ip 192.168.1.1 | jq .
{
  "ips": [
    {
      "available": "No",
      "ip": "192.168.1.1",
      "device": "some-network-name.domain.com",
      "label": "GigabitEthernet0/0/1"
    }
  ]
}
```

By default the raw JSON is received:
```
./d42cli get device sample.device.name
{"last_updated": "2018-06-13T14:43:07.064Z", "orientation": 1, "ip_addresses": [{"subnet": "DC1 1.1.4.0/23", "macaddress": "58:f3:9c:c3:e6:08", "subnet_id": 96, "ip": "1.1.4.19", "label": "eth0", "type": 1}], "serial_no": "SERIAL", "asset_no": "ASSETTAG", "rack": "Cage 12- Server", "manufacturer": "ACME Inc", "osver": "8", "device_purchase_line_items": [], "cpucore": "", "where": 5, "device_id": 2214}
```
