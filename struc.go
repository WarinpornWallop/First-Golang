package main

import "fmt"

//ออกแบบ structure
type Product struct{
	name string
	price float64
	category string
	discount int
}



func main(){
	product1 := Product{name: "ปากกา", price: 50.5, category: "เครื่องเขียน", discount: 10}
	product2 := Product{name: "เมาส์", price: 500, category: "อุปกรณ์คอม", discount: 20}
	
	fmt.Println(product1)
	fmt.Println(product2)
	product1.price = 100
	fmt.Println(product1.name)
	fmt.Println(product1.price)
}