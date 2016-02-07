import sys


def fizzbuzz(n):
    if n % 5 == 0 and n % 3 == 0:
        print 'fizzbuzz'
    elif n % 3 == 0:
        print 'fizz'
    elif n % 5 == 0:
        print 'buzz'
    else:
        print n


if __name__ == "__main__":
    x = int(sys.argv[1]) + 1
    for n in range(1,x):
        fizzbuzz(n)

