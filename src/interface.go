package main

import(
	"fmt"
)
//type empty interface{}
//type USB interface{
//	Name() string
//	Connect()
//}
//type Connecter interface{
//	Connect()
//}
//type PhoneConnecter struct{
//	name string
//}
//func (pc PhoneConnecter) Name() string{
//	return pc.name
//}
//func (pc PhoneConnecter) Connect(){
//	fmt.Println("Connectd:",pc.name)
//}
//func main(){

//	a := PhoneConnecter{"PhoneConnecter"}
//	a.Connect()//Connectd: PhoneConnecter
//	Disconnect(a)
//}
//func Disconnect(usb USB){
//	if pc,ok := usb.(PhoneConnecter);ok {
//		fmt.Println("Disconnected" ,pc.name)//Disconnected PhoneConnecter
//		return
//	}
//	fmt.Print("Disconnected.")//Disconnected.
//}

//type empty interface{}
//type USB interface{
//	Name() string
//	Connect()
//}
//type Connecter interface{
//	Connect()
//}
//type PhoneConnecter struct{
//	name string
//}
//func (pc PhoneConnecter) Name() string{
//	return pc.name
//}
//func (pc PhoneConnecter) Connect(){
//	fmt.Println("Connectd:",pc.name)
//}
//func main(){
//	a := PhoneConnecter{"PhoneConnecter"}
//	a.Connect()//Connectd: PhoneConnecter
//	Disconnect(a)
//}
//func Disconnect(usb interface{}){
//	switch v:=usb.(type){
//		case PhoneConnecter:
//		fmt.Println("Disconnected" ,v.name)
//		default:
//		fmt.Print("Unknown decive")
//	}
//}

