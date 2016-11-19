package main

func main() {

	num := 0
	num2 := 1

	result := 0
	for index := 0; index < 90; index++ {
		result = num + num2
		num = num2
		num2 = result
	}
	println(result)

}
