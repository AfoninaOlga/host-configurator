package main

import (
	"context"
	"flag"
	servicepb "github.com/AfoninaOlga/hostname-configurator/gen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func ParseFlag() string {
	var configPath string
	flag.StringVar(&configPath, "c", "configs/gateway.yaml", "Path to configuration file")
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
	viper.SetDefault("grpc_port", "8080")
	viper.SetDefault("grpc_addr", "127.0.0.1")
	viper.SetDefault("grpc_gw_port", "8090")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		viper.GetString("grpc_addr")+":"+viper.GetString("grpc_port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = servicepb.RegisterConfiguratorHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":" + viper.GetString("grpc_gw_port"),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%s\n", viper.GetString("grpc_gw_port"))
	log.Fatalln(gwServer.ListenAndServe())
}
