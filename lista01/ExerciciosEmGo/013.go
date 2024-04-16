package main
import("fmt")

func main (){
	var temp float32
	fmt.Scan(&temp)
	if temp>=9{
		fmt.Printf("NOTA = %f , CONCEITO = A ", temp)
	}else[
		if temp>=7.5{
			fmt.Printf("NOTA = %f , CONCEITO = B ", temp)
		}else{
			if temp>=6{
				fmt.Printf("NOTA = %f , CONCEITO = C ", temp)
			}else{
				fmt.Printf("NOTA = %f , CONCEITO = D ", temp)
			}
		}
	]
}
