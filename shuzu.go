package main
import "fmt"

func main(){
	strArr := [3]int{1,2,3}
	fmt.Printf("strArr[0] == %d\n",strArr[0])
	fmt.Printf("&strArr[0] == %p\n",&strArr[0])

	fmt.Printf("strArr[1] == %d\n",strArr[1])
	fmt.Printf("&strArr[1] == %p\n",&strArr[1])
	
	fmt.Printf("strArr[2] == %d\n",strArr[2])
	fmt.Printf("&strArr[2] == %p\n",&strArr[2])

	fmt.Printf("&strArr == %p\n",&strArr)
	newStrArr := new([3]int)
	newStrArr[0] = 1
	newStrArr[1] = 2 
	newStrArr[2] = 3 

	fmt.Printf("newstrArr[0] == %d\n",newStrArr[0])
	fmt.Printf("&newstrArr[0] == %p\n",&newStrArr[0])

	fmt.Printf("newstrArr[1] == %d\n",newStrArr[1])
	fmt.Printf("&newstrArr[1] == %p\n",&newStrArr[1])
	
	fmt.Printf("newstrArr[2] == %d\n",newStrArr[2])
	fmt.Printf("&newstrArr[2] == %p\n",&newStrArr[2])

	fmt.Println("strArr",strArr)
	fmt.Println("newStrArr",newStrArr)

	fmt.Printf("strArr == %p\n",&strArr)
	fmt.Printf("newStrArr == %p\n",newStrArr)

}
