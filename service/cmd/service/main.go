package main

import (
	"flag"
	servicepb "github.com/AfoninaOlga/hostname-configurator/gen"
	host_conf "github.com/AfoninaOlga/hostname-configurator/pkg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ParseFlag() string {
	var configPath string
	flag.StringVar(&configPath, "c", "configs/service.yaml", "Path to configuration file")
	flag.Parse()
	return configPath
}

func main() {
	// Parse command line flag
	configPath := ParseFlag()

	// Initialize and load configuration from file
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Set config default values
	viper.SetDefault("hostname_path", "/etc/hostname")
	viper.SetDefault("resolve_path", "/etc/resolv.conf")
	viper.SetDefault("grpc_port", "8080")

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":"+viper.GetString("grpc_port"))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Configurator service to the server
	servicepb.RegisterConfiguratorServer(s, host_conf.NewServer(viper.GetString("hostname_path"), viper.GetString("resolve_path")))
	// Serve gRPC Server
	log.Printf("Serving gRPC on 0.0.0.0:%s\n", viper.GetString("grpc_port"))
	log.Fatalln(s.Serve(lis))
}
