import re

board = {}
row_ends = {}
col_ends = {}
cur = None
facing = 0

data = open("input", "r").read().split("\n\n")

r = 0
for inp in data[0].split("\n"):
    r += 1
    for c, v in enumerate(inp, 1):
        if v != " ":
            board[r, c] = True if v == "." else False
            if board.get((r - 1, c)) == None:
                col_ends[c] = (r, None)
            else:
                col_ends[c] = (col_ends[c][0], r)
            if not cur and r == 1:
                cur = (1, c)
            row_ends[r] = (row_ends.get(r, (c,))[0], c)

path = tuple(re.findall("[0-9]+|[LR]", data[1]))

for cmd in path:
    if cmd.isdigit():
        steps = int(cmd)
        dp, di = not facing in (1, 3), (1, -1)[facing in (2, 3)]
        for _ in range(steps):
            new = cur[dp] + di
            ends = (col_ends, row_ends)[dp][cur[~dp]]
            if new < ends[0] or new > ends[1]:
                new = ends[0] if di == 1 else ends[1]
            if not board[(new_pos := (cur[~dp], new) if dp else (new, cur[~dp]))]:
                break
            cur = new_pos
    elif cmd == "R":
        facing = (1, 2, 3, 0)[facing]
    else:
        facing = (3, 0, 1, 2)[facing]

print(1000 * cur[0] + 4 * cur[1] + facing)