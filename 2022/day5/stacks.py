from collections import deque


class stacks:

    def __init__(self, lines: list[str]):
        size = len(lines[0]) // 4 + 1
        self._stacks = [deque() for i in range(size)]

        lines.pop()
        lines.reverse()
        for l in lines:
            self._push_line(l)

    def _push_line(self, line : str) -> None:
        for i in range(len(self._stacks)):
            letter = line[1+4*i]
            if letter != " ":
                self._stacks[i].append(letter)

    def pick_one_by_one(self, line: str) -> None:
        toks = line.split()
        num, fr, to = int(toks[1]), int(toks[3])-1, int(toks[5])-1
        for i in range(num):
            self._stacks[to].append(self._stacks[fr].pop())

    def pick_all(self, line: str) -> None:
        toks = line.split()
        num, fr, to = int(toks[1]), int(toks[3])-1, int(toks[5])-1
        tempd = deque()
        for i in range(num):
            tempd.append(self._stacks[fr].pop())
        for i in range(num):
            self._stacks[to].append(tempd.pop())

    def get_tops(self) -> list[str]:
        out = deque()
        for s in self._stacks:
            out.append(s[len(s)-1])
        return out
