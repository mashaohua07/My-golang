package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("phone start")
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

func (p Phone) Call() {
	fmt.Println("I am call")
}

type Camer struct {
	nos string
}

func (c Camer) Start() {
	fmt.Println("camer start")
}

func (c Camer) Stop() {
	fmt.Println("camer stop")
}

type Computer struct {

}

func (c Computer) Working(usb Usb) {
	usb.Start()
	//类型断言
	if phone, nk := usb.(Phone); nk {
		phone.Call()
	}
	usb.Stop()
}

func main()  {
	//computer := Computer{}
	//phone := Phone{}

	//computer.Working(phone)
	////computer.Working()


	var usbArr [3]Usb
	usbArr[0] = Phone{"view"}
	usbArr[1] = Phone{"bili"}
	usbArr[2] = Camer{"int"}

	computer := Computer{}
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
