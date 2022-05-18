package get

import (
	"pvectl/cli"
	"pvectl/proxmox"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get shows the current configuration an item in proxmox",
}

func init() {
	cli.RootCmd.AddCommand(getCmd)
}

func GetConfig(args []string, IDtype string) (err error) {
	id := cli.ValidateIDset(args, 0, IDtype+"ID")
	c := cli.NewClient()
	var config interface{}
	switch IDtype {
	case "MetricServer":
		config, err = proxmox.NewConfigMetricsFromApi(id, c)
	case "User":
		config, err = proxmox.NewConfigUserFromApi(id, c)
	}
	if err != nil {
		return
	}
	cli.PrintFormattedJson(getCmd.OutOrStdout(), config)
	return
}
