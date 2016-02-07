import sys

def main():
    
    with open(sys.argv[1]) as f:
        for line in f:
            solve(line)


def solve(line):
    (word, end) = line.split(',')
    letter = end.strip()

    pos = -1
    for x in range(0, len(word)):
        if word[x] == letter:
            pos = x

    print(pos)



if __name__ == '__main__':
    main()



