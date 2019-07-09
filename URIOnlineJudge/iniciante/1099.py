'''Leia um valor inteiro N que é a quantidade de casos de teste que vem a seguir. Cada caso de teste consiste de dois inteiros X e Y. Você deve apresentar a soma de todos os ímpares existentes entre X e Y.

Entrada
A primeira linha de entrada é um inteiro N que é a quantidade de casos de teste que vem a seguir. Cada caso de teste consiste em uma linha contendo dois inteiros X e Y.

Saída
Imprima a soma de todos valores ímpares entre X e Y.'''

N = int(input())
x = 0
y = 0
for i in range(N):
    X,Y = list(map(int,input().split()))
    soma = 0

    if(X==Y):
        print(0)
    else:
        tmp = 0
        if(X > Y):
            tmp = X
            X = Y
            Y = tmp

        while(X < (Y-1)):
            X += 1
            
            if(X%2 != 0):
                soma += X
                
        print(soma)
