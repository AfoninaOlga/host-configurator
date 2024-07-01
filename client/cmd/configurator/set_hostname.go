package configurator

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setHostname = &cobra.Command{
	Use:   "set-hostname [HOSTNAME]",
	Short: "Set hostname",
	Long:  "Command updates hostname in /etc/hostname",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]
		err := configurator.SetHostname(hostname)
		if err != nil {
			fmt.Printf("Couldn't update hostname. Reason: %v\n", err)
			return
		}
		fmt.Printf("Hostname is set to '%s'\n", hostname)
	},
}

func init() {
	rootCmd.AddCommand(setHostname)
}
