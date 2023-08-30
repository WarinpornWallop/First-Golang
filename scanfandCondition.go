package main

import "fmt" //นำเอา package เข้ามาทำงาน
func main(){
	// var name string
	// var score int 
	// fmt.Print("ป้อนชื่อนักเรียน = ")
	// fmt.Scanf("%s", &name)
	// fmt.Print("ป้อนคะแนนนักเรียน = ")
	// fmt.Scanf("%d", &score)
	// // fmt.Println("สวัสดี = ", name)
	// fmt.Println("คะแนน = ", score)

	// if score >= 50{
	// 	fmt.Println("สอบผ่าน")
	// }else{
	// 	fmt.Println("สอบไม่ผ่าน")
	// }
	var number int
	fmt.Print("ป้อนตัวเลช = ")		
	fmt.Scanf("%d", &number)
	//switch
	switch number{
	case 1:
		fmt.Println("เปิดบัญชีใหม่")	
	case 2:
		fmt.Println("ฝากเงิน - ถอนเงิน")	
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")	
	}
}