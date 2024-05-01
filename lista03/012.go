package main
import("fmt")


func mergesort(lista []int){
	n := len(lista)
	if n<2{
		return
	}
	mid := n/2
	var Llista []int
	var Rlista []int
	for c:=0;c<mid;c++{
		Llista = append(Llista, lista[c])
	}
	for c:=mid;c<n;c++{
		Rlista = append(Rlista, lista[c])
	}
	fmt.Print(Llista,"\n")
	fmt.Print(Rlista,"\n")
}

func main(){
	var x,temp int
	var reg []int
	fmt.Scan(&x)
	for c:=0;c<x;c++{
		fmt.Scan(&temp)
		reg = append(reg,temp)
	}
	mergesort(reg)
}