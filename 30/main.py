import sys

def main():
    with open(sys.argv[1], 'r') as test_cases:
        for test in test_cases:
            solve(test)


def solve(test):
    if test == '\n':
        return
    (first, second) = test.split(';')
    first_list = first.split(',')
    second_list = second.split(',')
    first_num_list = [int(x) for x in first_list]
    second_num_list = [int(x) for x in second_list]

    intersection = [x for x in first_num_list if x in second_num_list]

    for x in str(intersection):
        if not x.isdigit() and x != ',': 
            continue
        sys.stdout.write(str(x))

    sys.stdout.write('\n')


if __name__ == '__main__':
    main()
