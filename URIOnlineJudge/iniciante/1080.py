'''Leia 100 valores inteiros. Apresente então o maior valor lido e a posição dentre os 100 valores lidos.

Entrada
O arquivo de entrada contém 100 números inteiros, positivos e distintos.

Saída
Apresente o maior valor lido e a posição de entrada, conforme exemplo abaixo.'''

maior = 0
lista = {}
posicao = 0
while posicao < 100:
    valor = int(input())
    if(valor > maior):
        maior = valor
        lista[valor] = posicao
    posicao += 1
print(maior)
print(lista[maior]+1)
