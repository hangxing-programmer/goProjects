package basic

import "fmt"

// 定义一个接口
type Usb interface {
	Start()
	End()
}

type Phone struct {
	name string
}
type Camera struct {
	name string
}
type Computer struct {
}

// 让Phone实现Usb接口
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) End() {
	fmt.Println("手机结束工作...")
}

// Camera
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) End() {
	fmt.Println("相机结束工作...")
}

// 编写方法,只要实现了Usb接口(实现了Usb接口声明的所有方法)
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.End()
}

func main10() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	//实现多态
	var usbs [3]Usb
	usbs[0] = Phone{"OPPO"}
	usbs[1] = Phone{"VIVO"}
	usbs[2] = Camera{"xxxx"}
	fmt.Println(usbs)
}
