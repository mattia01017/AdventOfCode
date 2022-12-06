from stacks import stacks

try:
    line = "x"
    lines = []
    while True:
        line = input()
        if line == "":
            break
        lines.append(line)

    stacks = stacks(lines)

    while True:
        line = input()
        stacks.send_command1(line)

except EOFError:
    print("".join(stacks.get_tops()))