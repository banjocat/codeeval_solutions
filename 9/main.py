import sys

def main():
    with open(sys.argv[1]) as lines:
        for line in lines:
            solve(line)

def solve(line):
    print " ".join(every_other(line.strip().split(' ')))


class Stack:
    def __init__(self):
        self.stack = []
    
    def push(self, value):
        self.stack.append(value)


    def pop(self):
        value = self.stack[len(self.stack) - 1]
        self.stack = self.stack[:-1]
        return value


def every_other(line):
    stack = Stack()

    for c in line:
        stack.push(c)

    result = []
    for c in line:
        try:
            result.append(stack.pop())
            stack.pop()
        except:
            return result

    return result
    


if __name__ == '__main__':
    main()
