package main
import "fmt"

func main(){

	d := []int{1,2,3,4}
	e := d[2:]
	fmt.Println("d === ",d)
	fmt.Println("e === ",e)
	e[1] = 5
	fmt.Println("aftern e[1] = 5 e === ",e)
	fmt.Println("daftern e[1] = 5   === ",d)

	t := make([]int ,len(d),(cap(d)+1)*2)
	//for i := range d{
	//	t[i] = d[i]
	//}
	copy(t,d)
    d = t
	fmt.Println("t  === ",t)
	fmt.Println("d  === ",d)
}

