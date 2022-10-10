width = 50
init = [0] * (width-1) + [1] + [0] * (width-1)
def getLine(number):
    if number <= 1:
        return init
    prev = getLine(number-1)
    line = [prev[1]] + [int(prev[i-1] != prev[i+1]) for i in range(1, len(prev)-1)] + [prev[-2]]
    return line
for y in range(1, width+1):
    [print("X" if i else " ", end="") for i in getLine(y)]
    print("")