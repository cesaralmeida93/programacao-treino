'''Leia 5 valores Inteiros. A seguir mostre quantos valores digitados foram pares, quantos valores digitados foram ímpares, quantos valores digitados foram positivos e quantos valores digitados foram negativos.

Entrada
O arquivo de entrada contém 5 valores inteiros quaisquer.

Saída
Imprima a mensagem conforme o exemplo fornecido, uma mensagem por linha, não esquecendo o final de linha após cada uma.

'''

lista = []
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))
lista.append(int(input()))

par = 0
impar = 0
positivo = 0
negativo = 0

for x in lista:
    if x % 2 == 0:
        par += 1
    else:
        impar += 1

    if x > 0:
        positivo = positivo + 1
    elif x < 0:
        negativo = negativo + 1 

print("%d valor(es) par(es)"%par)
print("%d valor(es) impar(es)"%impar)
print("%d valor(es) positivo(s)"%positivo)
print("%d valor(es) negativo(s)"%negativo)
