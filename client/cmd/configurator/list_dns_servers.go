package configurator

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listServers = &cobra.Command{
	Use:   "dns-servers-list",
	Short: "List DNS servers",
	Long:  "Command lists DNS servers from /etc/resolv.conf",
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := configurator.ListServers()
		if err != nil {
			fmt.Printf("Couldn't get servers list. Reason: %v\n", err)
			return
		}
		if len(servers) == 0 {
			fmt.Println("No servers found.")
			return
		}
		fmt.Println("DNS servers:")
		for _, s := range servers {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(listServers)
}
