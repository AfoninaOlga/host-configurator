package configurator

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addServer = &cobra.Command{
	Use:   "dns-servers-add",
	Short: "Add DNS server",
	Long:  "Command adds DNS server to /etc/resolv.conf",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := args[0]
		err := configurator.AddServer(server)
		if err != nil {
			fmt.Printf("Couldn't add server. Reason: %v\n", err)
			return
		}
		fmt.Printf("Server %s added", server)
	},
}

func init() {
	rootCmd.AddCommand(addServer)
}
