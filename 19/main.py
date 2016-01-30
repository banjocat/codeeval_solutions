import sys
import re


def main():
    with open(sys.argv[1], 'r') as test_cases:
        for test in test_cases:
            split = str.split(test, ',')
            if split == ['\n']:
                continue
   
            bits = "{0:b}".format(int(split[0]))
            index1, index2 = (int(split[1]) * -1, int(split[2]) * -1)
            value1, value2 = (bits[index1], bits[index2])
            if value1 == value2:
                print('true')
            else:
                print('false')




if __name__ == '__main__':
    main()
