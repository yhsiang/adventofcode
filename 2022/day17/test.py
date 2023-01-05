import timeit

rocks = [[(0,0), (1,0), (2,0), (3,0)],        # horiz line
         [(0,1), (1,1), (2,1), (1,0), (1,2)], # plus
         [(0,0), (1,0), (2,0), (2,1), (2,2)], # backwards L
         [(0,0), (0,1), (0,2), (0,3)],        # vert line
         [(0,0), (0,1), (1,0), (1,1)]]        # square

def get_rock_width(rock):
    width = 0
    for coord in rock:
        width = max(width, coord[0]) 
    return width + 1

def check_collision(tower, rock, rockleft, rockbottom, rockwidth):
    if rockbottom < 0:
        return True
    if rockleft < 0 or rockleft + rockwidth > 7:
        return True
    for coord in rock:
        if (coord[0] + rockleft, coord[1] + rockbottom) in tower:
            return True
    return False

# surface map to use along with rock type and wind position to find a repeating cycle
def get_surface_map(tower, towerheight):
    surface = []
    for x in range(0, 7):
        y = towerheight
        while (x, y) not in tower and y > 0:
            y -= 1
        surface.append(towerheight - y)
    return tuple(surface)

def play_rocktris(wind, max_rocks):
    tower = set()
    rockcount = 0
    towerheight = 0
    windpos = 0
    newrock = 0
    height_adjust = 0
    savedstates = {}
    while rockcount < max_rocks:
        rock = rocks[newrock]
        w = get_rock_width(rock)
        rockbottom = towerheight + 3 
        rockleft = 2
        moving = True
        while moving:
            # move rock. First, jet of gas
            oldrockleft = rockleft
            if wind[windpos] == ">":
                rockleft += 1
            elif rockleft > 0:
                rockleft -= 1
            if check_collision(tower, rock, rockleft, rockbottom, w): #whoops
                rockleft = oldrockleft
            windpos = (windpos + 1) % len(wind)
            # then fall
            rockbottom -= 1
            if check_collision(tower, rock, rockleft, rockbottom, w):
                rockbottom += 1
                moving = False
                for coord in rock:
                    tower.add((coord[0] + rockleft, coord[1] + rockbottom))
                    towerheight = max(towerheight, coord[1] + rockbottom + 1)

        # fancy state matching code
        currstate = (get_surface_map(tower, towerheight), windpos, newrock)
        if currstate in savedstates:  # Hey, this looks familiar... 
            last_rockcount, last_height = savedstates[currstate]
            height_diff = towerheight - last_height
            count_diff = rockcount - last_rockcount
            repeats = (max_rocks - rockcount) // count_diff
            height_adjust += height_diff * repeats
            rockcount += count_diff * repeats       # skip ahead a bit...
        savedstates[currstate] = (rockcount, towerheight)

        newrock = (newrock + 1) % 5
        rockcount += 1
    return towerheight + height_adjust

def run():
    wind = open('input', 'r').read().rstrip()

    # Part 1
    print("pt 1:", play_rocktris(wind, 2022) )

    # Part 2
    print("pt 2:", play_rocktris(wind, 1000000000000) )

if __name__ == "__main__":
    print("\nDay 17: Pyroclastic Flow")
    starttime = timeit.default_timer()
    run()
    print("elapsed time: %fms\n" % ((timeit.default_timer()-starttime)*1000))