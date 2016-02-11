import sys

def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line.strip())



def solve(line):
    """
    make string all lowercase
    loop through each letter
    count occurence of letter place in map 
    sort values in map to a sorted list
    count down from 26 with product of list down till end
    """
    lower = line.lower()
    number_dict = {}
    for c in lower:
        if not c.isalpha():
            continue
        number_dict[c] = lower.count(c)

    letter_count = []
    for key in number_dict:
        letter_count.append(number_dict[key])

    happy_list = sorted(letter_count)[::-1]

    start = 26
    happy_sum = 0
    for num in happy_list:
        happy_sum += start * num
        start -= 1

    print happy_sum





if __name__ == '__main__':
    main()
