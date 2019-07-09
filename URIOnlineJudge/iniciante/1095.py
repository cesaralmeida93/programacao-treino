'''Você deve fazer um programa que apresente a sequencia conforme o exemplo abaixo.

Entrada
Não há nenhuma entrada neste problema.

Saída
Imprima a sequencia conforme exemplo abaixo'''

i, j = 1, 60

while(True):
    print("I=%d J=%d"%(i,j))
    i += 3
    j -= 5
    if j < 0:
        break
