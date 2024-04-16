package main
import("fmt")

func main(){
	var a,b,c int 
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Prinf("O VALOR DO DELTA E = ",(b*b)-(4*a*c))
}
