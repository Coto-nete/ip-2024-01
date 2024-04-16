package main
import("fmt")

func main(){
	var matricula int 
	ftm.Scan(&matricula)
	var temp,mp,ml,mt float32
	for c:=0;c<8;c++{
		fmt.Scan(&temp)
		mp +=temp
	}
	mp /=8
	for c:=0;c<5;c++{
		fmt.Scan(&temp)
		ml +=temp
	}
	ml /=5
	fmt.Scan(&temp)
	nt += temp
	var pres float32
	fmt.Scan(&pres)
	if (0.7*mp + 0.15*ml +0.15*nt>6){
		if pres>75/100{
			fmt.printf("Aprovado")
		}else{
			fmt.printf("Reprovado por frequencia ")
		}
	}else{
		if pres>75/100{
			fmt.printf("Reprovado por nota ")
		}else{
			fmt.printf("Reprovado por nota e frequencia ")
		}
	}
}
