'''A empresa ABC resolveu conceder um aumento de salários a seus funcionários de acordo com a tabela abaixo:


Salário	Percentual de Reajuste
0 - 400.00
400.01 - 800.00
800.01 - 1200.00
1200.01 - 2000.00
Acima de 2000.00

15%
12%
10%
7%
4%

Leia o salário do funcionário e calcule e mostre o novo salário, bem como o valor de reajuste ganho e o índice reajustado, em percentual.

Entrada
A entrada contém apenas um valor de ponto flutuante, com duas casas decimais.

Saída
Imprima 3 linhas na saída: o novo salário, o valor ganho de reajuste e o percentual de reajuste ganho, conforme exemplo abaixo.'''

a= float(input())
if (0<a<=400):
    rea=(((a*15)/100)+a)
    perc= 15
    print("Novo salario: %.2f" %rea)
    print("Reajuste ganho: %.2f" %(rea-a))
    print("Em percentual: " + str(perc) + " %")
elif(400<a<=800):
    rea=(((a*12)/100)+a)
    perc= 12
    print("Novo salario: %.2f" %rea)
    print("Reajuste ganho: %.2f" %(rea-a))
    print("Em percentual: " + str(perc) + " %")
elif(800<a<=1200):
    rea=(((a*10)/100)+a)
    perc= 10
    print("Novo salario: %.2f" %rea)
    print("Reajuste ganho: %.2f" %(rea-a))
    print("Em percentual: " + str(perc) + " %")
elif(1200<a<=2000):
    rea=(((a*7)/100)+a)
    perc= 7
    print("Novo salario: %.2f" %rea)
    print("Reajuste ganho: %.2f" %(rea-a))
    print("Em percentual: " + str(perc) + " %")
else:
    rea=(((a*4)/100)+a)
    perc= 4
    print("Novo salario: %.2f" %rea)
    print("Reajuste ganho: %.2f" %(rea-a))
    print("Em percentual: " + str(perc) + " %")
