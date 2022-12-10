n = 10

T = [(0, 0)] * n

touchedLast = list()
touchedLast.append(T[0])

touchedNearHead = list()
touchedNearHead.append(T[0])

move = {
    'U': lambda y: (y[0], y[1]+1),
    'D': lambda y: (y[0], y[1]-1),
    'R': lambda y: (y[0]+1, y[1]),
    'L': lambda y: (y[0]-1, y[1]),
}


def print_grid(T: list[list], size: int):
    for i in range(size//2, -size//2, -1):
        for j in range(-size//2, size//2):
            if (j, i) in T:
                print(T.index((j, i)), end='')
            else:
                print('.', end='')
        print()
    print('\n\n--------------------------------------\n\n')


for line in open(0):
    d, s = line[0], int(line[2:])
    for _ in range(s):
        next_move = move[d](T[0])
        T[0] = next_move
        for i in range(1, n):
            touch_x = T[i-1][0]-1 <= T[i][0] <= T[i-1][0]+1
            touch_y = T[i-1][1]-1 <= T[i][1] <= T[i-1][1]+1
            touch = touch_x and touch_y

            if not touch:
                diff = (T[i-1][0] - T[i][0],
                        T[i-1][1] - T[i][1])
                next_move = T[i-1]
                if diff[0] == 2 or diff[0] == -2:
                    next_move = (next_move[0] - diff[0]//2, next_move[1])
                if diff[1] == 2 or diff[1] == -2:
                    next_move = (next_move[0], next_move[1] - diff[1]//2)
                T[i] = next_move
        touchedLast.append(T[9])
        touchedNearHead.append(T[1])
    #print_grid(T, 20)

print(f"First part: {len(set(touchedNearHead))}")
print(f"Second part: {len(set(touchedLast))}")

