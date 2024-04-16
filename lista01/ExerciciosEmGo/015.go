package main 
import("fmt")

func main(){
	var temp int
	fmt.Scan(&temp)
	for c:=0;c<temp;c+=2{
		fmt.Printf("%i^2=%i",c,c*c)
	}
}
