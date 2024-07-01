package configurator

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getHostname = &cobra.Command{
	Use:   "get-hostname",
	Short: "Get hostname",
	Long:  "Command gets hostname from /etc/hostname",
	Run: func(cmd *cobra.Command, args []string) {
		hostname, err := configurator.GetHostname()
		if err != nil {
			fmt.Printf("Couldn't get hostname. Reason: %v\n", err)
			return
		}
		fmt.Printf("Hostname is '%s'\n", hostname)
	},
}

func init() {
	rootCmd.AddCommand(getHostname)
}
