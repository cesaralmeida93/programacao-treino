'''Leia 3 valores de ponto flutuante A, B e C e ordene-os em ordem decrescente, de modo que o lado A representa o maior dos 3 lados. A seguir, determine o tipo de triângulo que estes três lados formam, com base nos seguintes casos, sempre escrevendo uma mensagem adequada:
se A ≥ B+C, apresente a mensagem: NAO FORMA TRIANGULO
se A2 = B2 + C2, apresente a mensagem: TRIANGULO RETANGULO
se A2 > B2 + C2, apresente a mensagem: TRIANGULO OBTUSANGULO
se A2 < B2 + C2, apresente a mensagem: TRIANGULO ACUTANGULO
se os três lados forem iguais, apresente a mensagem: TRIANGULO EQUILATERO
se apenas dois dos lados forem iguais, apresente a mensagem: TRIANGULO ISOSCELES
Entrada
A entrada contem três valores de ponto flutuante de dupla precisão A (0 < A) , B (0 < B) e C (0 < C).

Saída
Imprima todas as classificações do triângulo especificado na entrada.'''

# -*- coding: utf-8 -*-

a,b,c = map(float, input().split())
if(a>b>c)or(a==b>c)or(a==b>c) or(a==b==c):
    A=a
    B=b
    C=c
elif(a>c>b)or(a==c>b)or(a>c==b):
    A=a
    B=c
    C=b
elif(b>a>c)or(b==a>c)or(b>a==c):
    A=b
    B=a
    C=c
elif(b>c>a)or(b==c>a)or(b>c==a):
    A=b
    B=c
    C=a
elif(c>a>b)or(c==a>b)or(c>a==b):
    A=c
    B=a
    C=b
elif(c>b>a)or(c==b>a)or(c>b==a):
    A=c
    B=b
    C=a

if(A>=B+C):
    print("NAO FORMA TRIANGULO")
else:
    if(A**2 == B**2+C**2):
        print("TRIANGULO RETANGULO")
    if(A**2>B**2+C**2):
        print("TRIANGULO OBTUSANGULO")
    if(A**2<B**2+C**2):
        print("TRIANGULO ACUTANGULO")
    if(A==B==C):
        print("TRIANGULO EQUILATERO")
    elif(A==B):
        print("TRIANGULO ISOSCELES")
    elif(A==C):
        print("TRIANGULO ISOSCELES")
    elif(B==C):
        print("TRIANGULO ISOSCELES")
