import sys
from math import sqrt

f = open(sys.argv[1])
data = f.read().split()
f.close()


def clean(x):
    ret = float(x[:-2])
    if 'ms' not in x:
        ret *= 0.001
    return ret


data = [clean(x) for x in data]

mean = sum(data) / len(data)
std = sqrt(sum([(x-mean)**2 for x in data])/len(data))

print(mean)
print(std)
