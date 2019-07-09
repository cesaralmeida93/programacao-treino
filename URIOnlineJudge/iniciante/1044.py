'''Leia 2 valores inteiros (A e B). Após, o programa deve mostrar uma mensagem "Sao Multiplos" ou "Nao sao Multiplos", indicando se os valores lidos são múltiplos entre si.

Entrada
A entrada contém valores inteiros.

Saída
A saída deve conter uma das mensagens conforme descrito acima.'''

# -*- coding: utf-8 -*-

def mult():
    v1, v2 = [float(i) for i in input().split()]
    if v2 % v1 == 0 or v1 % v2 == 0:
        print("Sao Multiplos")
    else:
        print("Nao sao Multiplos")
mult()
