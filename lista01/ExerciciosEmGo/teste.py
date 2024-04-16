for c in range(int(input())):
    linha = list(map(float,input().split(" ")))
    print("A RENDA DO JOGO ",c+1," E = ", linha[0]*linha[1]/100 + (linha[0]*linha[2]/100)*5 + (linha[0]*linha[3]/100)*10 + (linha[0]*linha[4]/100)*20)