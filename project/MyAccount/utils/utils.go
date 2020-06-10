package utils

import (
	"fmt"
)

type Account struct {

	key string
	loop bool

	//账户初始余额
	balance float64
	//每次收支金额
	money float64
	//每次收支说明
	note string
	//定义一个变量,记录是否有收支行为
	flag bool
	//收支详情  进行拼接处理
	details string
}

func (this *Account) ShowDetails() {
	fmt.Println("===========当前收支明细==========")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支情况， ...来一笔吧！")
	}
}

func NewAccount() *Account {
	return &Account{
		"",
		true,
		10000,
		0.0,
		"",
		false,
		"收支\t账户金额\t收支金额\t说明\t",
	}
}

func (this *Account) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v\t%v\t%v", this.balance, "", this.money, "", this.note)
	this.flag = true
}

func (this *Account) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("支出余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v\t%v\t%v", this.balance, "", -this.money, "", this.note)
	this.flag = true
}

func (this *Account) quit() {
	fmt.Println("你确定要退出么? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请从新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *Account) MainMenu() {
	for  {
		fmt.Println("\n=============收支记账软件============")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Print("请选择1--4:")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.ShowDetails()

		case "2":
			this.income()

		case "3":
			this.pay()

		case "4":
			this.quit()

		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出程序")
}
