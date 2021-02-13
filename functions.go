package main
import "fmt"

func main() {
defer hello()
a := add(5, 4)
fmt.Println("x + y = ",a)
b := multiply(4,5)
fmt.Println(b)

var f func(x,y int) int = add
fmt.Println(f(3,5))

// анонимные функции
fun := func(x,y int) int {return x + y}
fmt.Println(fun(5,15));

}

func hello(){
	fmt.Println("Hello")
}

func add(x int, y int) int{
	z:= x + y

	return z
}

func multiply(x int, y int) (int){
	f:= x * y
	return f
}




