package main
import("fmt")

func main (){
	var temp float 
	fmt.Scan(&temp)
	if temp <=300{
		fmt.Printf("SALARIO COM REAJUSTE : %f",temp*1.5)
	}else{
		fmt.Printf("SALARIO COM REAJUSTE : %f",temp*1.3)
	}
}
