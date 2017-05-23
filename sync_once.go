package main
import "time"
import "fmt"
import "sync"

var a string
var once sync.Once

func setup(){
	a = "hello , world"
	fmt.Println("set string : ",a)
}

func doprint(){
	once.Do(setup)
	fmt.Println("string : ",a)
}
func main(){
	go doprint()
	go doprint()
	time.Sleep(time.Second)
}
