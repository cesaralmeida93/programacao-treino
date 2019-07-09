'''Leia um valor inteiro. A seguir, calcule o menor número de notas possíveis (cédulas) no qual o valor pode ser decomposto. As notas consideradas são de 100, 50, 20, 10, 5, 2 e 1. A seguir mostre o valor lido e a relação de notas necessárias.

Entrada
O arquivo de entrada contém um valor inteiro N (0 < N < 1000000).

Saída
Imprima o valor lido e, em seguida, a quantidade mínima de notas de cada tipo necessárias, conforme o exemplo fornecido. Não esqueça de imprimir o fim de linha após cada linha, caso contrário seu programa apresentará a mensagem: “Presentation Error”.'''

# -*- coding: utf-8 -*-

numero = int(input())
if 0 < numero < 1000000:
    print(numero)
    x = numero // 100
    print('{} nota(s) de R$ 100,00'.format(x))
    y = numero % 100
    z = y // 50
    print('{} nota(s) de R$ 50,00'.format(z))
    w = y % 50
    q = w // 20
    print('{} nota(s) de R$ 20,00'.format(q))
    u = w % 20
    i = u // 10
    print('{} nota(s) de R$ 10,00'.format(i))
    v = u % 10
    n = v // 5
    print('{} nota(s) de R$ 5,00'.format(n))
    j = v % 5
    h = j // 2
    print('{} nota(s) de R$ 2,00'.format(h))
    a = j % 2
    b = a // 1
    print('{} nota(s) de R$ 1,00'.format(b))
