package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	showMessage("Tonnam")
	total(50, 100)
	total2(50, 100)
	delivery := getDelivery()
	fmt.Println("ค่าจัดส่ง = ",delivery)
	myCart := getTotalCart(500, 500)
	fmt.Println("ยอดรวมทั้งหมด = ", myCart)
}
func showMessage(name string) {
	fmt.Println("Hello!", name)
}
func total(number1 int, number2 int){
	fmt.Println("ยอดรวม = ", number1+number2)
}
func total2(number1, number2 int){
	fmt.Println("ยอดรวม = ", number1+number2)
}
func getDelivery() int { //int คือประเภทของค่าที่ return
	return 50
}
func getTotalCart(number1, number2 int) int {
	total := number1 + number2
	return total
}