package main
import ("fmt"
)

func main(){
<<<<<<< Updated upstream
<<<<<<< HEAD
	var in,fm int 
	var reg []int
	
}
=======
    for{
        var t,f int
        var ent,resp string
        fmt.Scan(&t)
        fmt.Scan(&f)
        if t==0{
            break
        }
        fmt.Scan(&ent)
        for c:=t-1;c>=0;c--{
            cont:=0
            if len(resp)<f{
                resp = string(ent[c:])
            }else{
                if string(ent[c])>string(resp[cont]){
                    resp = string(ent[c])+string(resp[cont+1:])
                }
            }
        }
        fmt.Print(resp,"\n")
        
    }
}
>>>>>>> a647f183765cbf582532cfb5ab9d382083777731
=======
	for{
		var in,fm int 
		var ent,st string
		fmt.Scan(&in)
		fmt.Scan(&fm)
		if in==0{
			break
		}
		fmt.Scan(&ent)
		st = ent[in-fm:]
		for c:=in-fm-1;c>=0;c--{
			for j:=0;j<fm;j++{
				if int(st[j])<int(ent[c]){
					
				}
			}
		}
	}
}
>>>>>>> Stashed changes
