package main
import "fmt"
import "time"
func main(){
	jobs := make(chan int ,5)
	done := make(chan bool)

	go func(){
		for{
			j,more := <-jobs
			if more {
				fmt.Println("received job",j)
			}else{
				fmt.Println("received all jobs")
				time.Sleep(time.Second*5)
				done <- true
				return 
			}
			
		}
	}()

	for j:=1; j<=30; j++{
		jobs <- j
		fmt.Println("sent job")
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<- done
}
