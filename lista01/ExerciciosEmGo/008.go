package main
import("fmt")

func main(){
	var temp1, temp2 int 
	fmt.Scan(&temp1)
	fmt.Scan(&temp2)
	fmt.Printf("O VALOR DE CUSTO E = %.2f", (2*(3.14159*(temp1*temp1))+(2*3.14159*temp1*temp2))*100)
}
