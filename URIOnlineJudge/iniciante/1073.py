'''Leia um valor inteiro N. Apresente o quadrado de cada um dos valores pares, de 1 até N, inclusive N, se for o caso.

Entrada
A entrada contém um valor inteiro N (5 < N < 2000).

Saída
Imprima o quadrado de cada um dos valores pares, de 1 até N, conforme o exemplo abaixo.

Tome cuidado! Algumas linguagens tem por padrão apresentarem como saída 1e+006 ao invés de 1000000 o que ocasionará resposta errada. Neste caso, configure a precisão adequadamente para que isso não ocorra.'''

n = int(input())
a = 2
pot = 0

if n > 5 and n < 2000:
    if n % 2 == 0:
        while a <= n:
            pot = a ** 2
            print("%d^2 = %d"%(a,pot))
            a += 2
    
    else:
        while a < n:
            pot = a ** 2
            print("%d^2 = %d"%(a,pot))
            a += 2

