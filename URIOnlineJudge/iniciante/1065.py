'''
Faça um programa que leia 5 valores inteiros. Conte quantos destes valores digitados são pares e mostre esta informação.

Entrada
O arquivo de entrada contém 5 valores inteiros quaisquer.

Saída
Imprima a mensagem conforme o exemplo fornecido, indicando a quantidade de valores pares lidos.'''

lista = []
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))

par = 0

for x in lista:
    if x % 2 == 0:
        par += 1

print("%d valores pares"%par)
