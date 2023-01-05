with open("input") as f:
    obsidians = {(int(x), int(y), int(z)) for x, y, z in [c.split(',') for c in f.read().split('\n')]}

neighbours = [(1, 0, 0), (-1, 0, 0), (0, 1, 0), (0, -1, 0), (0, 0, 1), (0, 0, -1)]
visible_sides = sum([1 for x, y, z in obsidians for nx, ny, nz in neighbours if (x + nx, y + ny, z + nz) not in obsidians])

print(visible_sides)