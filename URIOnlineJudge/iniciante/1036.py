'''Leia 3 valores de ponto flutuante e efetue o cálculo das raízes da equação de Bhaskara. Se não for possível calcular as raízes, mostre a mensagem correspondente “Impossivel calcular”, caso haja uma divisão por 0 ou raiz de numero negativo.

Entrada
Leia três valores de ponto flutuante (double) A, B e C.

Saída
Se não houver possibilidade de calcular as raízes, apresente a mensagem "Impossivel calcular". Caso contrário, imprima o resultado das raízes com 5 dígitos após o ponto, com uma mensagem correspondente conforme exemplo abaixo. Imprima sempre o final de linha após cada mensagem.'''

# -*- coding: utf-8 -*-

import math
a, b, c = [float(i) for i in input().split()]
x=(b**2)-(4*a*c)
if  x<0:
   print("Impossivel calcular")
elif a==0:
  print("Impossivel calcular")        
else: 
 x = math.sqrt(x)
 x1=(-b + x) / (a*2)
 x2=(-b - x) / (a*2)
 print("R1 = %.5f" %x1)
 print("R2 = %.5f" %x2)
