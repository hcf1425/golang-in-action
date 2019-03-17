package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected to :", pc.name)
}

func main() {
	a := PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)
}

//ok pattern
//func Disconnect(usb USB) {
//	if pc, ok := usb.(PhoneConnector); ok {
//		fmt.Println("Disconnected", pc.name)
//		return
//	}
//	fmt.Println("Unknown device!")
//}

//switch type  ---> 传入空接口时用得比较多。
func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Diconnected", v.name)
	default:
		fmt.Println("Unknown device!")
	}

}
