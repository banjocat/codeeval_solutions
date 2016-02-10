import sys


def main():
    with open(sys.argv[1]) as f:
        for line in f:
            solve(line)


def are_you_happy(num):
    past = []

    while num not in past:
        past.append(num)
        num = happy_number(num)
        if num == "1":
            print 1
            return

    print "0"



def solve(line):
    are_you_happy(line.strip())


def happy_number(num_str):
    num = 0
    for d in num_str:
        num += int(d) * int(d)

    return str(num)
    


if __name__ == '__main__':
    main()
