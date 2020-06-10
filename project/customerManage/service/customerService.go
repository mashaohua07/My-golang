package service

import (
	"src/project/customerManage/model"
)

type CustomerService struct {
	customer []model.Customer

	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

func NewCustomerService() *CustomerService  {
	customerServer := &CustomerService{}
	customerServer.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "mashaohua07@sina.com")
	customerServer.customer = append(customerServer.customer, customer)
	return customerServer
}

func (this *CustomerService) List() []model.Customer {
	return this.customer
}

