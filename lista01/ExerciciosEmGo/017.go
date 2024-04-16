package main 
import("fmt")

func main (){
	var x,y int 
	fmt.Scan(&x)
	fmt.Scan(&y)
	if x%2==0{
		for c:=x;c>x+y;c+=2{
			Printf("%i ",c)
		}
	}else{
		fmt.Printf("O PRIMEIRO NUMERO NAO E PAR.")
	}
}
