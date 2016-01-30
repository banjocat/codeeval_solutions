import sys

def main():
    with open(sys.argv[1], 'r') as test_cases:
        for test in test_cases:
            solve(test)


def solve(test):
   split = str.split(test, ',')
   if split == ['\n']:
       return

   uniques = sorted({int(x) for x in split})
   uniques_str = [str(x) for x in uniques]
   print(','.join(uniques_str))


if __name__ == '__main__':
    main()
