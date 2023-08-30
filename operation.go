package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	// var number1 = 10
	// var number2 = 3
	number1, number2 := 10,3
	var number3, number4 = 20,6
	fmt.Println("ผลบวก = ", number1+number2 )
	fmt.Println("ผลลบ = ", number1-number2 )
	fmt.Println("ผลคูณ = ", number3*number4 )
	fmt.Println("ผลหาร = ", number1/number2 )
	fmt.Println("เศษ = ", number3%number4 )
	fmt.Println("เท่ากันหรือไม่ = ", number3==number4 )
	fmt.Println("ไม่เท่ากันหรือไม่ = ", number3!=number4 )
	fmt.Println(number3, "มากกว่า ", number2, "=" , number3 > number2 )
	fmt.Println(number3, "มากกว่าหรือเท่ากับ", number2, "=" , number3 >= number2 )
}