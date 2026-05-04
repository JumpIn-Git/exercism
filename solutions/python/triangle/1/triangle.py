def valid(sides):
    return all(c > 0 for c in sides) and (sides[0] + sides[1] >= sides[2]) and (sides[1] + sides[2] >= sides[0]) and (sides[0] + sides[2] >= sides[1])

def equilateral(sides):
    return valid(sides) and len(set(sides)) == 1


def isosceles(sides):
    return valid(sides) and len(set(sides)) == 2 or len(set(sides)) == 1

1
def scalene(sides):
    return valid(sides) and len(set(sides)) == 3
