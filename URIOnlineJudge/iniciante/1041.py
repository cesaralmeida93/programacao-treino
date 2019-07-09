'''Leia 2 valores com uma casa decimal (x e y), que devem representar as coordenadas de um ponto em um plano. A seguir, determine qual o quadrante ao qual pertence o ponto, ou se está sobre um dos eixos cartesianos ou na origem (x = y = 0).



Se o ponto estiver na origem, escreva a mensagem “Origem”.

Se o ponto estiver sobre um dos eixos escreva “Eixo X” ou “Eixo Y”, conforme for a situação.

Entrada
A entrada contem as coordenadas de um ponto.

Saída
A saída deve apresentar o quadrante em que o ponto se encontra.'''

# -*- coding: utf-8 -*-

def coord():
    p = input().split()
    X = float(p[0])
    Y = float(p[1])

    if (X == 0.0) and (Y == 0.0):
        print("Origem")
    elif (X == 0.0) and (Y != 0):
        print("Eixo Y")
    elif (X != 0.0) and (Y == 0):
        print("Eixo X")
    elif (X > 0.0) and (Y > 0):
        print("Q1")
    elif (X < 0.0) and (Y > 0):
        print("Q2")
    elif (X < 0.0) and (Y < 0):
        print("Q3")
    else:
        print("Q4")

coord()
