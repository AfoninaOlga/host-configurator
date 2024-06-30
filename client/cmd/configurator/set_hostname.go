package configurator

import (
	"github.com/spf13/cobra"
	"log"
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
			log.Printf("error updating hostname: %v\n", err)
			return
		}
		log.Printf("Hostname set to %s\n", hostname)
	},
}

func init() {
	rootCmd.AddCommand(setHostname)
}
