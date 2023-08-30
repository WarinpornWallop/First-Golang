package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	var name string = "Tonnamwarineiei"
	var age int = 25
	var score float32 = 95.8
	var isPass bool = true
	nickname := "tonnam"
	name = "Tonnameiei"
	const fixname string = "konsuay"
	fmt.Println("Hello Go Programming")
	fmt.Println("My name is ", name)
	fmt.Println("My Nickname is ", nickname)
	fmt.Println("My Realname is ", fixname)
	fmt.Println("Age = ", age)
	fmt.Println("Score = ", score)
	fmt.Println("Pass Exam = ", isPass)
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Type name is %T \n", name)
	fmt.Printf("Age = %v \n", age)
	fmt.Printf("Type Age is %T \n", age)
	//fmt.Println("Tonnamwarineiei")
}