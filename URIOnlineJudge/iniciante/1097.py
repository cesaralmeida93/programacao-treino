'''Você deve fazer um programa que apresente a sequencia conforme o exemplo abaixo.

Entrada
Não há nenhuma entrada neste problema.

Saída
Imprima a sequencia conforme exemplo abaixo.'''

# -*- coding: utf-8 -*-
I=-1
J=5
for i in range(5):
    i=I+2
    j=J+2
    print("I=%d J=%d"%(i,j))
    print("I=%d J=%d"%(i,j-1))
    print("I=%d J=%d"%(i,j-2))
    I=i
    J=j
