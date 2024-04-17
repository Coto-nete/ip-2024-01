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
	var n int
	cont :=1
	fmt.Scan(&n)
	for c:=0;c<n-1;c++{
		for j:=c+1;j<n;j++{
			fmt.Printf("Final %v: Time%v X Time%v",cont,c+1,j+1)
			cont++
		}
	}
}
