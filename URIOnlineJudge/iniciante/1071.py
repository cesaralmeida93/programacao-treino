'''Leia 2 valores inteiros X e Y. A seguir, calcule e mostre a soma dos números impares entre eles.

Entrada
O arquivo de entrada contém dois valores inteiros.

Saída
O programa deve imprimir um valor inteiro. Este valor é a soma dos valores ímpares que estão entre os valores fornecidos na entrada que deverá caber em um inteiro.'''

x = int(input())
y = int(input())

soma = 0

if x > y:
    troca = x
    x = y
    y = troca

if x % 2 != 0:
    x = x + 2
    while x < y:
        if x % 2 != 0:
            soma = soma + x
        x += 1

else:
    x = x + 1
    while x < y:
        if x % 2 != 0:
            soma = soma + x
        x = x + 1

print(soma)
