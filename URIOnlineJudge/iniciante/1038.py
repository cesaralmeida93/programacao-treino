'''Com base na tabela abaixo, escreva um programa que leia o código de um item e a quantidade deste item. A seguir, calcule e mostre o valor da conta a pagar.



Entrada
O arquivo de entrada contém dois valores inteiros correspondentes ao código e à quantidade de um item conforme tabela acima.

Saída
O arquivo de saída deve conter a mensagem "Total: R$ " seguido pelo valor a ser pago, com 2 casas após o ponto decimal.
'''

# -*- coding: utf-8 -*-

a, b = [int(i) for i in input().split()]
if a == 1:
    r = 4.00 * b
elif a == 2:
    r = 4.50 * b
elif a == 3:
    r = 5.00 * b
elif a == 4:
    r = 2.00 * b
else:
    r = 1.50 * b
print("Total: R$ %.2f"%r)
