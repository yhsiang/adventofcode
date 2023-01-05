with open("input") as f:
    obsidians = {(int(x), int(y), int(z)) for x, y, z in [c.split(',') for c in f.read().split('\n')]}

visible_sides = 0
neighbours = [(1, 0, 0), (-1, 0, 0), (0, 1, 0), (0, -1, 0), (0, 0, 1), (0, 0, -1)]
visited = set()
queue = [(0, 0, 0)]

while queue:
    x, y, z = queue.pop(0)
    visited.add((x, y, z))
    for nx, ny, nz in neighbours:
        check = (x + nx, y + ny, z + nz)
        if -1 <= check[0] <= 22 and -1 <= check[1] <= 22 and -1 <= check[2] <= 22 and check not in visited and check not in queue:
            if check in obsidians:
                visible_sides += 1
            else:
                queue.append(check)

print(visible_sides)