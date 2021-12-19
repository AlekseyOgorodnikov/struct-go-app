package main

import (
	"log"
	"struct-go/internal/ports"

	"struct-go/internal/adapters/app/api"
	"struct-go/internal/adapters/core/arithmetic"

	gRPC "struct-go/internal/adapters/framework/left/grpc"
	"struct-go/internal/adapters/framework/right/db"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	dbDriver := "mysql"
	dsourceName := "root:root@tcp(127.0.0.1:3306)/hex_test"

	dbaseAdapter, err = db.NewAdapter(dbDriver, dsourceName)
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
