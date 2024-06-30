package configurator

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "configurator",
	Short: "hostname configurator CLI client managing /etc/hostname and /etc/resolv.conf",
	Long: `
───────╔═╦╗────────╔╗─────
╔═╦═╦═╦╣═╬╬═╦╦╦╦╦═╗║╚╦═╦╦╗
║═╣╬║║║║╔╣║╬║║║╔╣╬╚╣╔╣╬║╔╝
╚═╩═╩╩═╩╝╚╬╗╠═╩╝╚══╩═╩═╩╝
──────────╚═╝─────────────
Change hostname; list, add or remove DNS servers.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

//var configurator = conf.NewConfigurator("http://localhost:8090")
