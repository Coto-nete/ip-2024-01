package main
import("fmt")

func main(){
	var i,f int
	var ent,temp string 
	for {
		fmt.Scan(&i)
		fmt.Scan(&f)
		if i ==0 ||f==0 {
			break
		}
		fmt.Scan(&ent)
		for c:=i-f;c>=i;c++{
			temp += ent[c]
		}
		fmt.Print(temp)
	}
}
