package main

import (
	"log"
	"os"
	"struct-go/internal/ports"

	"struct-go/internal/adapters/app/api"
	"struct-go/internal/adapters/core/arithmetic"

	gRPC "struct-go/internal/adapters/framework/left/grpc"
	"struct-go/internal/adapters/framework/right/db"
)

func main() {
	var err error
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DS_NAME", "root:Root147258*@tcp(127.0.0.1:3306)/hex_test")

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	// dbaseDriver := "mysql"
	// dsourceName := "root:root@tcp(127.0.0.1:3306)/hex_test" or root:root@tcp(db:3306)/hex_test

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}

//CloceDcConnection
