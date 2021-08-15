package main

import (
	"fmt"

	///"github.com/kataras/iris"
	"log"
	"net/http"
)

//FUNCTION MAIN
func main() {
	Port:= "5555"
	http.Handlefunc("/",ServeFiles)
	fmt.Println("Serving @ : ","127.0.0.1"+Port)
	ListenAndServe(Port,nil)
// //FIBONACCI

	// 	var fibonancciValue int = 8
	// 	fmt.Println(fibonacci(fibonancciValue))

	// //CALCULATOR

	// 	fmt.Println(calculator("+",0,1))
	// 	fmt.Println(calculator("-",0,1))
	// 	fmt.Println(calculator("*",10,7))
	// 	fmt.Println(calculator("/",1,0))

	// //READ AND WRITE FILE TEXT

	// 	// content, err:=ioutil.ReadFile("./src/test/read.txt");
	// 	// if err != nil {
	// 	// 	fmt.Println(err)
	// 	// }
	// 	// fmt.Println(string(content))

	// 	// content := []byte("12345")
	// 	// err = ioutil.WriteFile("./src/test/write.txt",content,0644)
	// 	// if err!= nil {
	// 	// 	return
	// 	// }

	// //READ FILE AND CONV TO INT
	// 	content, err:=ioutil.ReadFile("./src/test/read.txt");
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	num:= strings.Split(string(content)," ")
	// 	var number []int
	// 	for i:=0;i<len(num);i++{
	// 		byteToInt, _ := strconv.Atoi(string(num[i]))
	// 		number = append(number,byteToInt)
	// 	}
	// 	fmt.Println("Chuỗi nhận được từ file")
	// 	printArr(number)
	// 	number2:=number
	// //SẮP XẾP LỚN DẦN
	// 	for i:=0;i<len(number);i++{
	// 		for j:=i;j<len(number);j++{
	// 			if(number[j]<=number[i]){
	// 				temp:= number[i]
	// 				number[i]= number[j]
	// 				number[j] = temp
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("Chuỗi sắp xếp lớn dần")
	// 	printArr(number)

	// //SẮP XẾP NHỎ DẦN
	// 	for i:=0;i<len(number);i++{
	// 		for j:=i;j<len(number);j++{
	// 			if(number[j]>=number[i]){
	// 				temp:= number[i]
	// 				number[i]= number[j]
	// 				number[j] = temp
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("Chuỗi sắp xếp nhỏ dần")
	// 	printArr(number)
	// //tìm kiếm
	// 	fmt.Printf("Số nhỏ nhất dãy là %d\n", searchMin(number))
	// 	fmt.Printf("Số lớn nhất dãy là %d\n", searchMax(number))
	// 	fmt.Printf("Số trung bình là %d\n", searchMedium(number))
	// // buble sort
	// 	for i:=0;i<len(number);i++{
	// 		for j:=0;j<len(number)-i-1;j++{
	// 			if number[j]>number[j+1]{
	// 				temp:= number[j]
	// 				number[j] = number[j+1]
	// 				number[j+1] = temp

	// 			}

	// 		}
	// 	}
	// 	fmt.Println("dãy số dau khi buble sort là ")
	// 	printArr(number)
	// //kiểm tra số nguyên tố
	// 	for i:=0;i<len(number2);i++{
	// 		if kiemtrasonguyento(number2[i]){
	// 			fmt.Printf("%d là số nguyên tố\n",number2[i])
	// 		}else{
	// 			fmt.Printf("%d không số nguyên tố\n",number2[i])
	// 		}
	// 	}
	// //tim kiem gia tri trong file
	// 	x := 11
	// 	if search(number2,x){
	// 		fmt.Printf("giá trị %d tồn tại\n", x)
	// 	}else{
	// 		fmt.Printf("giá trị %d không tồn tại\n", x)
	// 	}
}
//FUNCTION
func ServeFiles(w http.ResponseWriter, r *http.Request){
	switch r.method{
	case "GET":
		path = r.URL.Path
		fmt.Println(path)
		if path =="/"{
			path= "./html/index.html"
		}else {
			path = "."+path
		}
	case "POST":
		r.ParseMultipartForm(0)
	}
}
//dãy fibonacci
func fibonacci(n int) int {
	var value int = 2
	var1, var2 := 1, 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i < n; i++ {
			//fmt.Println(var1,var2)
			value = var1 + var2
			var1 = var2
			var2 = value
		}
	}
	return value
}

//phép tính
func calculator(math string, num1 float32, num2 float32) float32 {
	switch math {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		if num2 == 0 {
			return 0
		}
		return num1 / num2
	case "*":
		return num1 * num2
	default:
		return 0
	}
}

//in chuỗi
func printArr(number []int) {
	for i := 0; i < 13; i++ {
		fmt.Printf("%d ", number[i])
	}
	fmt.Println()
}

//tim kiem gia tri nhỏ nhat
func searchMin(number []int) int {
	min := number[0]
	for i := 0; i < len(number); i++ {
		if number[i] < min {
			min = number[i]
		}
	}
	return min
}

//tim kiem gia tri lon nhat
func searchMax(number []int) int {
	max := number[0]
	for i := 0; i < len(number); i++ {
		if number[i] > max {
			max = number[i]
		}
	}
	return max
}

//tinh trung binh day ra kieu int
func searchMedium(number []int) int {
	tong := 0
	x := len(number)
	for i := 0; i < x; i++ {
		tong += number[i]
	}
	return int(tong / x)
}

// kiểm tra só nguyên tố
func kiemtrasonguyento(x int) bool {
	flag := true
	count := 0
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			count++
		}
	}
	if count != 0 {
		flag = false
	}
	return flag
}

//tim kiem gia tri cho trươc
func search(number []int, x int) bool {
	flag := false
	for i := 0; i < len(number); i++ {
		if x == number[i] {
			flag = true
		}
	}
	return flag
}
