'''Leia 3 valores reais (A, B e C) e verifique se eles formam ou não um triângulo. Em caso positivo, calcule o perímetro do triângulo e apresente a mensagem:


Perimetro = XX.X


Em caso negativo, calcule a área do trapézio que tem A e B como base e C como altura, mostrando a mensagem


Area = XX.X

Entrada
A entrada contém três valores reais.

Saída
O resultado deve ser apresentado com uma casa decimal.'''

# -*- coding: utf-8 -*-

a, b, c = [float(i) for i in input().split()]
if abs(b-c) < a and a < b + c:
    print("Perimetro = %.1f"%(a+b+c))
elif abs(a-c) < b and b < a + c:
    print("Perimetro = %.1f"%(a+b+c))
elif abs(a-b) < c and c < a + b:
    print("Perimetro = %.1f"%(a+b+c))
else:
    print("Area = %.1f"%(c*(a+b)/2))
