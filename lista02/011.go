package main 
import("fmt")


func r(n float64)float64{
	if n==0{
		return 0
	}else{
		fmt.Printf( "%.2f",r(k-1)+(n/(r(k-1)))/2)
	}
}

func main(){
	var n float64
	r(n)
}