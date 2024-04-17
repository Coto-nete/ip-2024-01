package main
import("fmt")

func main (){
	var n,temp int
	fmt.Scan(&n)
	var nums []int
	var res int 
	for c:=0;c<n;c++{
		fmt.Scan(&temp)
		nums = append(nums, temp)
	}
	for c:=0;c<n-1;c++{
		achar := 1
		contar := 1
		for j:=c+1;j<n;j++{
			if nums[j]==nums[c]+achar{
				achar ++
				contar++
			}
		}
		if contar>=res{
			res = contar
		}
	}
	fmt.Printf("O comprimento do segmento crescente maximo e : %v\n",res)
}
