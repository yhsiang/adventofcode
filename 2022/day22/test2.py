import re

board = {}
row_ends = {}
col_ends = {}
sections = set()
cur = None
facing = 0
side = 50

data = open("input", "r").read().split("\n\n")

r = 0
for inp in data[0].split("\n"):
    row_section = r // side + 1
    r += 1
    for c, v in enumerate(inp, 1):
        col_section = (c - 1) // side + 1
        if v != " ":
            sections.add((row_section, col_section))
            board[r, c] = (
                (True, (row_section, col_section))
                if v == "."
                else (False, (row_section, col_section))
            )
            if board.get((r - 1, c)) == None:
                col_ends[c] = (r, None)
            else:
                col_ends[c] = (col_ends[c][0], r)
            if not cur and r == 1:
                cur = (1, c)
            row_ends[r] = (row_ends.get(r, (c,))[0], c)


def turn(facing, dir="R", times=1):
    return (0, 1, 2, 3)[(facing + times * (-1, 1)[dir == "R"]) % 4]


path = tuple(re.findall("[0-9]+|[LR]", data[1]))

for cmd in path:
    if cmd.isdigit():
        steps = int(cmd)
        for _ in range(steps):
            dp, di = not facing in (1, 3), (1, -1)[facing in (2, 3)]
            new = cur[dp] + di
            s1, s2 = board.get(cur)[1][:: -1 if dp else 1]
            ends = (col_ends, row_ends)[dp][cur[~dp]]
            if new < ends[0] or new > ends[1]:
                changed_sec_pos = 1 if di == -1 else side
                const_sec_pos = cur[~dp] - side * (s2 - 1)
                for sec, value in (
                    (
                        (s1 - 3 * di, s2),
                        (facing, (side + 1) - changed_sec_pos, const_sec_pos),
                    ),
                    (
                        (s1 - 3 * di, s2 + 2 * di),
                        (facing, (side + 1) - changed_sec_pos, const_sec_pos),
                    ),
                    (
                        (s1 - 3 * di, s2 + 2 * di),
                        (facing, (side + 1) - changed_sec_pos, const_sec_pos),
                    ),
                    (
                        (s1 + di, s2 - di),
                        (turn(facing, ("R", "L")[dp]), const_sec_pos, changed_sec_pos),
                    ),
                    (
                        (s1 + di, s2 + di),
                        (
                            turn(facing, ("L", "R")[dp]),
                            (side + 1) - const_sec_pos,
                            (side + 1) - changed_sec_pos,
                        ),
                    ),
                    (
                        (s1 - 3 * di, s2 - di),
                        (
                            turn(facing, ("L", "R")[dp]),
                            (side + 1) - changed_sec_pos,
                            (side + 1) - const_sec_pos,
                        ),
                    ),
                    (
                        (s1 - 3 * di, s2 + di),
                        (turn(facing, ("R", "L")[dp]), const_sec_pos, changed_sec_pos),
                    ),
                    (
                        (s1 - di, s2 - 3 * di),
                        (
                            turn(facing, ("L", "R")[dp]),
                            (side + 1) - changed_sec_pos,
                            const_sec_pos,
                        ),
                    ),
                    (
                        (s1 - di, s2 + 3 * di),
                        (turn(facing, ("R", "L")[dp]), const_sec_pos, changed_sec_pos),
                    ),
                    (
                        (s1 - di, s2 + 2 * di),
                        (
                            turn(facing, times=2),
                            changed_sec_pos,
                            (side + 1) - const_sec_pos,
                        ),
                    ),
                    (
                        (s1 - di, s2 - 2 * di),
                        (
                            turn(facing, times=2),
                            changed_sec_pos,
                            (side + 1) - const_sec_pos,
                        ),
                    ),
                    (
                        (s1 + di, s2 + 2 * di),
                        (
                            turn(facing, times=2),
                            changed_sec_pos,
                            (side + 1) - const_sec_pos,
                        ),
                    ),
                    (
                        (s1 + di, s2 - 2 * di),
                        (
                            turn(facing, times=2),
                            changed_sec_pos,
                            (side + 1) - const_sec_pos,
                        ),
                    ),
                ):
                    s = sec[:: -1 if dp else 1]
                    if s in sections:
                        new_facing = value[0]
                        new_pos = (
                            value[1] + side * (sec[0] - 1),
                            value[2] + side * (sec[1] - 1),
                        )
                        new_pos = new_pos[:: -1 if dp else 1]
                        break
            else:
                new_pos = (cur[~dp], new) if dp else (new, cur[~dp])
                new_facing = facing
            if not board[new_pos][0]:
                break
            cur = new_pos
            facing = new_facing
    else:
        facing = turn(facing, cmd)

print(1000 * cur[0] + 4 * cur[1] + facing)