package main
import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int 
}

func main(){

	var Book1 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 111
	
	fmt.Printf("Book1 title : %s \n",Book1.title)
	fmt.Printf("Book1 author : %s \n",Book1.author)
}
