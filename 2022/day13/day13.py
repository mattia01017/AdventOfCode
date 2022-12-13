from functools import cmp_to_key
from json import loads
from sys import stdin

def compare_packets(a: list, b: list):
    for i in range(min(len(a), len(b))):
        x = a[i]
        y = b[i]
        if type(a[i]) == type(b[i]) == int:
            if x != y:
                return x < y
        else:
            if type(x) == int:
                x = [x]
            if type(y) == int:
                y = [y]
            if x != y:
                return compare_packets(x, y)
    return len(a) <= len(b)


class packet_comparator(list):
    def __lt__(self, other):
        return compare_packets(self, other)


res = set()
packets = list()
i = 1

while True:
    try:
        l1 = input()
        if len(l1) == 0:
            continue
        l2 = input()

        p1 = loads(l1)
        p2 = loads(l2)
        packets.append(p1)
        packets.append(p2)
        if compare_packets(p1, p2):
            res.add(i)
        i += 1
    except EOFError:
        break

print("First part: " + str(sum(res)))

packets.append([[2]])
packets.append([[6]])

s_packets = sorted(packets, key=packet_comparator)

print(f"Second part: " +
      str((s_packets.index([[2]])+1) * (s_packets.index([[6]])+1)))
