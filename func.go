package main

import ( 
	"fmt"
	"math"
)
func sum(nums ... int){
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main(){
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)
	
}


