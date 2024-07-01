package configurator

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteServer = &cobra.Command{
	Use:   "dns-servers-delete",
	Short: "Add DNS server",
	Long:  "Command deletes DNS server from /etc/resolv.conf",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]
		err := configurator.DeleteServer(server)
		if err != nil {
			fmt.Printf("Couldn't delete server. Reason: %v\n", err)
			return
		}
		fmt.Printf("Server %s is deleted\n", server)
	},
}

func init() {
	rootCmd.AddCommand(deleteServer)
}
