'''Leia 3 valores inteiros e ordene-os em ordem crescente. No final, mostre os valores em ordem crescente, uma linha em branco e em seguida, os valores na sequência como foram lidos.

Entrada
A entrada contem três números inteiros.

Saída
Imprima a saída conforme foi especificado.'''

# -*- coding: utf-8 -*-

a, b, c = [int(i) for i in input().split()]

if a > b and a > c:
    if b > c:
        print("%d\n%d\n%d\n"%(c,b,a))
    else:
        print("%d\n%d\n%d\n"%(b,c,a))
elif a < b and a < c:
    if b > c:
        print("%d\n%d\n%d\n"%(a,c,b))
    else:
        print("%d\n%d\n%d\n"%(a,b,c))
elif a < b and a > c:
    print("%d\n%d\n%d\n"%(c,a,b))
else:
    print("%d\n%d\n%d\n"%(b,a,c))
print("%d\n%d\n%d"%(a,b,c))
