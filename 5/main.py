import sys
import exceptions



def main():
    with open(sys.argv[1]) as f:
        for line in f:
            solve(line.strip())



def find_first_sequence(numbers):
    """
    loop through each number:
        2nd loop starting at -1 of current:
            if not match continue:
            if matches:
                check if sequence added by both are equal
                    if not conitnue
                    else done
    """
    for i in range(0, len(numbers)):
        if i == 0:
            continue
        for j in range(0, i):
            if numbers[i] != numbers[j]:
                continue
            try:
                value = sequence_at(i, j, numbers)
                return value
            except:
                continue

def sequence_at(i, j, numbers):
    match = []
    for offset in range(0, len(numbers)):
        if numbers[i + offset] in match:
            return match
        if numbers[i + offset] != numbers[j + offset]:
            raise exceptions.ValueError
        match.append(numbers[i + offset])


def solve(line):
    line = line.split(' ')
    result = find_first_sequence(line) 
    print " ".join(result)






if __name__ == '__main__':
    main()
