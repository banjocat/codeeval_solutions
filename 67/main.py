import sys

def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line.strip())



def solve(line):
    """
    Could of used just python conversion
    but that would of been boring
    """
    hex_to_dec = {
            '0': 0,
            '1': 1,
            '2': 2,
            '3': 3,
            '4': 4,
            '5': 5,
            '6': 6,
            '7': 7,
            '8': 8,
            '9': 9,
            'a': 10,
            'b': 11,
            'c': 12,
            'd': 13,
            'e': 14,
            'f': 15
            }
    hex_number = line[::-1]
    integer = 0
    for x in range(0, len(hex_number)):
        integer += hex_to_dec[hex_number[x]] * 16 ** x

    print integer



if __name__ == '__main__':
    main()
