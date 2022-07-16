package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { //return type 명시
	return a * b
}
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words[1])
}

// defer => function이 끝나고 실행될 것, 자주 사용함
func lenAndUpperReturn(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func canIDrink(age int) bool {
	// variable 선언 가능, if else 구문에서만 사용하능 변수 선언(variable expression)
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func canIDrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// func main() {
// 	totalLength, upperCase := lenAndUpperReturn("hwalim")
// 	fmt.Println(totalLength, upperCase)
// }

// func main() {
// 	total := superAdd(1, 2, 3, 4, 5, 6)
// 	fmt.Println(total)
// }

// func main() {
// 	// fmt.Println(canIDrink(16))
// 	fmt.Println(canIDrinkSwitch(18))
// }

// 메모리 주소 확인하기 (&), 메모리 값 확인하기 (*)
// func main() {
// 	a := 2
// 	b := &a // 메모리 저장
// 	*b = 20
// 	fmt.Println(a)
// }

// func main() {
// 	names := []string{"nico", "lynn", "dal"}
// 	names = append(names, "alala")
// 	fmt.Println(names)
// }

// func main() {
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	for key, value := range nico {
// 		fmt.Println(key, value)
// 	}
// 	fmt.Println(nico)
// }

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kinchi", "ramen"}
	nico := person{name: "nico", age: 16, favFood: favFood}
	fmt.Println(nico)
}
