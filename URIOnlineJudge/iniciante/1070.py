'''
Leia um valor inteiro X. Em seguida apresente os 6 valores ímpares consecutivos a partir de X, um valor por linha, inclusive o X ser for o caso.

Entrada
A entrada será um valor inteiro positivo.

Saída
A saída será uma sequência de seis números ímpares.'''

x = int(input())

a = 0

while True:
    if(x%2 != 0):
        a += 1
        print(x)
    x += 1
    if(a == 6):
        break
