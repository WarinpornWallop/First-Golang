package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	
	numbers := []int{100,200,300}
	numbers = append(numbers, 400)
	numbers = append(numbers, 500)
	fmt.Println(numbers)
	fmt.Println(numbers[:])
	fmt.Println(numbers[1:]) //indexที่ 1 จนถึงสุดท้าย
	fmt.Println(numbers[1:3]) //indexที่ 1 จนถึง 3-1
	fmt.Println(numbers[:3])
}