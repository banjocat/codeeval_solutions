import sys
from itertools import filterfalse

def main():
    with open(sys.argv[1], 'r') as test_cases:
        for test in test_cases:
            solve(test)


if 

def solve(test):
    if test == '\n':
        return
    (first, second) = test.split(';')
    first_list = first.split(',')
    second_list = second.split(',')
    first_set = [int(x) for x in first_list]
    second_set = [int(x) for x in second_list]

    intersection = sorted(first_set.intersection(second_set))
    for x in filterfalse(lambda j: not j.isdigit() and j != ',', str(intersection)):
        sys.stdout.write(str(x))
    sys.stdout.write('\n')

if __name__ == '__main__':
    main()
