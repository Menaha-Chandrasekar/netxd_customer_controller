package main

import (
	"context"
	"net"

	"fmt"

	config "github.com/Menaha-Chandrasekar/netxd_config"
	constants "github.com/Menaha-Chandrasekar/netxd_constants"
	pro "github.com/Menaha-Chandrasekar/netxd_customer" // Import the generated Go code
	netxdcustomercontroller "github.com/Menaha-Chandrasekar/netxd_customer_controller/customer_controller"
	service "github.com/Menaha-Chandrasekar/netxd_dal/netxd_dal_service"
	tc "github.com/Menaha-Chandrasekar/netxd_transaction"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "bankgrpc", "profiles")
	rpc.CustomerService = service.InitCustomerService(customerCollection, context.Background())
}
func initTransaction(client *mongo.Client){
	rpc.TransactionService = service.InitTransaction(client, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	initTransaction(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &netxdcustomercontroller.RPCServer{})
    tc.RegisterTransactionServiceServer(s, &netxdcustomercontroller.TransactionSever{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}