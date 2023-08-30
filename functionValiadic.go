package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	result := summation2(101,200,21)
	fmt.Println("ผลรวม = ", result)
}
// ... คือรับกี่จำนวนก็ได้ ไม่ต้องมา fix parameter
func summation2(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}