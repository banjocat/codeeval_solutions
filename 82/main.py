import sys

def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line.strip())


def solve(line):
    armstrong = sum_of_nth_power_of_digits(line)
    print armstrong == line


def sum_of_nth_power_of_digits(number):
    total = 0
    power = len(number)
    for c in number:
        total += int(c) ** power

    return str(total)




if __name__ == '__main__':
    main()
