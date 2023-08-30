package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	result, check := summation(101,200)
	fmt.Println("ผลรวม = ", result)
	fmt.Println(check)
}
func summation(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0{
		status = "เลขคู่"
	}else{
		status = "เลขคี่"
	}
	return total, status
}