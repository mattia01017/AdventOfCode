H = (0,0)
T = (0,0)

touched = list()
touched.append(T)

directions = {
    'U' : lambda y,s: (y[0], y[1]+s),
    'D' : lambda y,s: (y[0], y[1]-s),
    'R' : lambda y,s: (y[0]+s, y[1]),
    'L' : lambda y,s: (y[0]-s, y[1]),
}

def exec_and_log(to_update : tuple, callable, logs : list):
    ret = callable(to_update)
    logs.append(ret)
    return ret

for line in open(0):
    d, s = line[0], int(line[2:])
    
    H = directions[d](H, s)
    
    while T[0] < H[0]-1:
        T = exec_and_log(T, lambda y: (y[0]+1, y[1]), touched)
    while T[0] > H[0]+1:
        T = exec_and_log(T, lambda y: (y[0]-1, y[1]), touched)
    while T[1] < H[1]-1:
        T = exec_and_log(T, lambda y: (y[0], y[1]+1), touched)
    while T[1] > H[1]+1:
        T = exec_and_log(T, lambda y: (y[0], y[1]-1), touched)
    T = H

print(len(set(touched))-1)