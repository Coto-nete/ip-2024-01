package main 
import("fmt")

func main (){
	var temp,i,r,e int 
	fmt.Scan(&i)
	fmt.Scan(&r)
	fmt.Scan(&e)
	for c:=i;c<i+(e-1)*r;c+=r{
		temp += c 
	}
	fmt.Printf("%i",temp)
}
