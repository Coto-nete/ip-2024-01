package main
import(
	"fmt"
	"math")

func main(){
	var x,c,resul,rg float64
	fmt.Scan(&x)
	fmt.Scan(&rg)
	den := 1
	resul = 1
	for c=1;c<rg;c++{
		for j:=2;j<=int(2*c);j++{
			den *=j
		}
		temp := (math.Pow(x,2*c))/float64(den)
		if int(c)%2==0{
			resul += temp
		} else{
			resul -= temp
		}
	}
	fmt.Print(resul,"\n")
	fmt.Printf("cos(%.2f) = %.6f",x,resul)
}