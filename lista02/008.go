package main
import("fmt")

func fac(n int)int{
	if n==1{
		return n 
	}else{
		return fac(n-1)*n
	}
}

func main(){
	var n,possiveis int
	var temp string
	var times []string
	fmt.Scan(&n)
	possiveis = (fac(n)/(2*fac(n-2)))
	for c:=0;c<n;c++{
		temp = "time"+string(c+1)
		fmt.Print(temp)
		times = append(times,temp)
	}
	fmt.Print(times)
	fmt.Print(possiveis)
}
