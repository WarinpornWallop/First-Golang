package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){

	// country := map[string]string{"TH":"ไทย","JP":"ญี่ปุ่น"}
	country := map[string]string{"TH":"ไทย","JP":"ญี่ปุ่น"}
	country["TH"] = "ไทย"
	country["JP"] = "ญี่ปุ่น"
	fmt.Println(country["JP"])
	value, check := country["JP"] //check=ตรวจสอบว่ามี key มั้ย ส่วน value คือเอาค่านั้นมา
	if check {
		fmt.Println(value)
	}else{
		fmt.Println("ไม่พบข้อมูล")
	}
}