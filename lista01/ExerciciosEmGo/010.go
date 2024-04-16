package main 
import("fmt")

func main (){
	var a,b,c,d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Printf("O VALOR DO DETERMINANTE E = ",(a*d)-(b*c))
}
