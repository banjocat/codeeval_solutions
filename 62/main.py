import sys


def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line.strip())


def solve(line):
    numbers = line.split(',')
    print int(numbers[0]) % int(numbers[1])



if __name__ == '__main__':
    main()

