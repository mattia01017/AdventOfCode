from stacks import stacks

try:
    line = "x"
    lines = []
    while True:
        line = input()
        if line == "":
            break
        lines.append(line)

    stacks1 = stacks(lines.copy())
    stacks2 = stacks(lines)

    while True:
        line = input()
        stacks1.pick_one_by_one(line)
        stacks2.pick_all(line)

except EOFError:
    print(f"pick-one-by-one stacks tops: " + "".join(stacks1.get_tops()))
    print(f"pick-all stacks tops: " + "".join(stacks2.get_tops()))