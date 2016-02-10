import sys

def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line)




def solve(line):
    print is_self_described_number(line.strip())


def is_self_described_number(num_str):
    """
    loop through each:
        count how many of that index there are in the string
            if == to value at index continue
            if != to value at index return "0"
    return "1"
    """
    for i in range(0, len(num_str)):
        expected = int(num_str[i])
        if num_str.count(str(i)) != expected:
            return "0"

    return "1"




if __name__ == '__main__':
    main()

