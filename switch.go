package main
import "fmt"
import "time"

func main(){
	//basic switch
	i := 2
	fmt.Print("Write ",i,"as ")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	//you can usd "," to separate multiple expressions
	switch time.Now().Weekday(){
		case time.Saturday,time.Sunday:
			fmt.Println("It is weekend")
		default:
			fmt.Println("It is weekday")
	}

	//switch without an expression is an alternate way to express
	// if/else logic
	t := time.Now()
	switch{
		case t.Hour() < 12:
			fmt.Println("It is before noon")
		default:
			fmt.Println("It is after noon")
	}

	// a type switch compare types instead of values
	 whatAmI := func(i interface{}){
		switch t := i.(type) {
			case bool:
				fmt.Println("I am bool")
			case int:
				fmt.Println("I am an int")
			default :
				fmt.Println("Do not know type %T\n",t)
		}
	 }
	 whatAmI(true)
	 whatAmI(1)
	 whatAmI("hey")

}
