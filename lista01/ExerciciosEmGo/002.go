package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var temp,renda int 
	fmt.Scan()
	for c:=0;c>temp;c++{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var nuns []int
		for _, r := range input {
			if r != ' ' {
				num, err := strconv.Atoi(string(r))
				if err != nil {
					fmt.Println("Error converting character to integer:", err)
					return
				}
				nuns = append(nuns, num)
			}
		}
		renda = nuns[0]*nuns[1]/100 + (nuns[0]*nuns[2]/100)*5 + (nuns[0]*nuns[3]/100)*10 + (nuns[0]*nuns[4]/100)*20
		fmt.Print("A renda do jogo N. ",c+1,"=",renda)
	}

}
