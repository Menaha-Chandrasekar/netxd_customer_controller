package main

import (
	"context"
	
	config "github.com/Menaha-Chandrasekar/netxd_config"
	constants "github.com/Menaha-Chandrasekar/netxd_constants"
	netxdcustomercontroller "github.com/Menaha-Chandrasekar/netxd_customer_controller/netxd_customer_controller"
	netxddalservice "github.com/Menaha-Chandrasekar/netxd_dal/netxd_dal_service"
	"net"

	"fmt"

	pro "github.com/Menaha-Chandrasekar/netxd_customer/customer_proto" // Import the generated Go code

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "bankgrpc", "profiles")
	netxdcustomercontroller.CustomerService = netxddalservice.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &netxdcustomercontroller.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}