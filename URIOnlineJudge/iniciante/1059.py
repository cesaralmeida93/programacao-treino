'''Faça um programa que mostre os números pares entre 1 e 100, inclusive.

Entrada
Neste problema extremamente simples de repetição não há entrada.

Saída
Imprima todos os números pares entre 1 e 100, inclusive se for o caso, um em cada linha.'''

i = 1
while(i <= 100):
    i = i + 1
    if i % 2 == 0:
        print(i)
    
