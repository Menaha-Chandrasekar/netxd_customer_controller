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


func(s *RPCServer)CreateCustomer(ctx context.Context,req * pro.CustomerRequest)(*pro.CustomerResponse,error){
	dbProfile:=&netxddalmodels.CustomerRequest{
		CustomerId: req.CustomerId,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		BankId:     req.BankId,
		Balance:    req.Balance,
	}
	res,err:=CustomerService.CreateCustomer(dbProfile)
	if err != nil {
		return nil, err
	}else {
		responseProfile := &pro.CustomerResponse{
			CustomerId: res.CustomerId,
			CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		return responseProfile, nil
	}

}