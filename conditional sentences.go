package main
import "fmt"

func main() {
	var a int = 5
	var	b int = 3
	var c bool = a == b
	fmt.Println(c)

	var numbers [5]int = [5]int{1,2,3,4,5}
	fmt.Println(numbers)
	numbers2 := [...]int{1,2,3,4}
	fmt.Println(numbers2)


	if a < b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	switch a {
	case 9: 
		fmt.Println("9")
	case 8:
		fmt.Println("8")
	default:
		fmt.Println("Default")
		
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i*i);
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for _, value := range users {
		fmt.Println(value);
	}
}
