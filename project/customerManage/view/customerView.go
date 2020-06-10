package main

import (
	"fmt"
	"src/project/customerManage/service"
)

type fdvbf interface {
	mainMenu()
}

type customerView struct {
	key string //接收用户输入
	loop bool  //表示是否循环的显示主菜单
	customerService *service.CustomerService
}

func (this *customerView) List() {
	customers := this.customerService.List()
	fmt.Println("--------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------客户列表完成---------------------")
}

func (c *customerView) mainMenu() {
	for  {
		fmt.Println("-------------------客户信息管理软件-------------------")
		fmt.Println("                   1 添 加 客 户")
		fmt.Println("                   2 修 改 客 户")
		fmt.Println("                   3 删 除 客 户")
		fmt.Println("                   4 客 户 列 表")
		fmt.Println("                   5  退    出")
		fmt.Println("请选择（1-5）: ")

		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			c.List()
		case "5":
			c.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !c.loop  {
			break
		}
	}

	fmt.Println("你已退出系统...")
}

func main()  {
	customerView := customerView{
		key: "",
		loop: true,
	}

	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
