package main
import("fmt")

func main (){
	var n int 
	var temp float32
	fmt.Scan(&n)
	for c:=1;c<=n;c++{
		temp=temp+1/c
	}
	fmt.Printf("&f",temp)
}
