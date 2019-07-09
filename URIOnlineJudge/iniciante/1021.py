'''Leia um valor de ponto flutuante com duas casas decimais. Este valor representa um valor monetário. A seguir, calcule o menor número de notas e moedas possíveis no qual o valor pode ser decomposto. As notas consideradas são de 100, 50, 20, 10, 5, 2. As moedas possíveis são de 1, 0.50, 0.25, 0.10, 0.05 e 0.01. A seguir mostre a relação de notas necessárias.

Entrada
O arquivo de entrada contém um valor de ponto flutuante N (0 ≤ N ≤ 1000000.00).

Saída
Imprima a quantidade mínima de notas e moedas necessárias para trocar o valor inicial, conforme exemplo fornecido.

Obs: Utilize ponto (.) para separar a parte decimal.'''

# -*- coding: utf-8 -*-

n=float(input()) 
while n>=1000000.00 or n<=0:
    n=input()
n += 0.001
c100=int(n/100)
n=n-c100*100
c50=int(n/50)
n=n-c50*50
c20=int(n/20)
n=n-c20*20
c10=int(n/10)
n=n-c10*10
c5=int(n/5)
n=n-c5*5
c2=int(n/2)
n=n-c2*2
m100=int(n)
n=n-m100
m50=int(n/0.5)
n=n-m50*0.5
m25=int(n/0.25)
n=n-m25*0.25
m10=int(n/0.10)
n=n-m10*0.10
m5=int(n/0.05)
n=n-m5*0.05
m1=int(n/0.01)
n=n-m1*0.01
print("NOTAS:")
print("%d nota(s) de R$ 100.00" % (c100))
print("%d nota(s) de R$ 50.00" % (c50))
print("%d nota(s) de R$ 20.00" % (c20))
print("%d nota(s) de R$ 10.00" % (c10))
print("%d nota(s) de R$ 5.00" % (c5))
print("%d nota(s) de R$ 2.00" % (c2))
print("MOEDAS:")
print("%d moeda(s) de R$ 1.00" % (m100))
print("%d moeda(s) de R$ 0.50" % (m50))
print("%d moeda(s) de R$ 0.25" % (m25))
print("%d moeda(s) de R$ 0.10" % (m10))
print("%d moeda(s) de R$ 0.05" % (m5))
print("%d moeda(s) de R$ 0.01" % (m1))



