package main
import "fmt"

func main() {

	var x int = 4
	px := &x
	fmt.Println(px)
	fmt.Println(*px)
	d:=5
	cangeValue(&d)
	fmt.Println(d)
	fmt.Println(add(4,5))
}

func cangeValue(x *int){
	*x = (*x) * (*x)
}