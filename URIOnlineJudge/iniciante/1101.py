'''Leia um conjunto não determinado de pares de valores M e N (parar quando algum dos valores for menor ou igual a zero). Para cada par lido, mostre a sequência do menor até o maior e a soma dos inteiros consecutivos entre eles (incluindo o N e M).

Entrada
O arquivo de entrada contém um número não determinado de valores M e N. A última linha de entrada vai conter um número nulo ou negativo.

Saída
Para cada dupla de valores, imprima a sequência do menor até o maior e a soma deles, conforme exemplo abaixo.'''

while True:
    try:
        m, n = list(map(int,input().split()))
        if(m < 1 or n < 1):
            break
        tmp = 0

        if(m > n):
            tmp = m
            m = n
            n = tmp
        soma = 0
        resultado = ''
        while(m <= n):
            resultado += str(m) + ' '
            soma += m
            m += 1
        resultado += 'Sum=%d'
        print(resultado %soma)
    except:
        break