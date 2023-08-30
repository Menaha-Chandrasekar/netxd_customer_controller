package netxdcustomercontroller

import (
	"context"
	netxddalinterface "module/netxd_dal/netxd_dal_interface"
	netxddalmodels "module/netxd_dal/netxd_dal_models"
	pro "module/netxd_customer/customer_proto"

)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxddalinterface.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerRequest) (*pro.CustomerResponse, error) {
	customerrequest := &netxddalmodels.CustomerRequest{CustomerId: req.CustomerId}
	result, err := CustomerService.CreateCustomer(customerrequest)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}