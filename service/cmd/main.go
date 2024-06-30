package main

import (
	"context"
	host_conf "github.com/AfoninaOlga/hostname-configurator"
	servicepb "github.com/AfoninaOlga/hostname-configurator/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

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
	viper.SetDefault("grpc_gw_port", "8090")
	viper.SetDefault("hostname_path", "/etc/hostname")
	viper.SetDefault("resolve_path", "/etc/resolv.conf")

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":"+viper.GetString("grpc_port"))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Configurator service to the server
	servicepb.RegisterConfiguratorServer(s, host_conf.NewServer(viper.GetString("hostname_path"), viper.GetString("resolve_path")))
	// Serve gRPC Server
	log.Printf("Serving gRPC on 0.0.0.0:%s\n", viper.GetString("grpc_port"))
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		"0.0.0.0:"+viper.GetString("grpc_port"),
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
