package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	for count :=1; count <= 10; count++{
	
		if count == 5{
			// break //เบรคแล้วไม่ทำต่อ
			continue //กระโดดข้ามค่านั้น ที่ไม่อยากให้ทำ
		}
		fmt.Println(count)
	}
	fmt.Println("จบโปรแกรม")
	numbers := []int{10,20,30,40,50,60,70,80,90,100}
	for index := 0; index < len(numbers); index++ {
		fmt.Println("Index = ",index, "Value = ", numbers[index])
		// fmt.Println(numbers[index])
	}

	//for range
	for index, value := range numbers{
		fmt.Println("Index = ",index, "Value = ",value)
	}
	language := map[string]string{"TH":"ไทย","JP":"ญี่ปุ่น","EN":"อังกฤษ"}
	for _, value := range language{ // _ คือ ไม่ต้องมี index ไม่ต้องสนใจ ignore
		fmt.Println( "Value = ",value)
	}
}