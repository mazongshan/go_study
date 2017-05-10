package main
import "fmt"

func main(){
	var map_c map[string]string
	map_c = make(map[string]string)

	map_c["F"] = "P"
	map_c["I"] = "R"

	for country := range map_c {
		fmt.Println("captial of ",country,"is",map_c[country])
	}
	
	delete(map_c,"F")

	for country := range map_c {
		fmt.Println("captial of ",country,"is",map_c[country])
	}
}
