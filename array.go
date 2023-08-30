package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	// var numbers [3] int
	// fmt.Println(numbers)
	// numbers[0] = 100
	// numbers[1] = 200
	// numbers[2] = 300
	var numbers [3] int = [3]int{100,200,300}
	numbers2 := [3]int{100,200,300}
	numbersnolimit := [...]int{100,200,300,400,500,600,700,800,900}
	size := len(numbers)
	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(numbersnolimit)
	fmt.Println(size)
}