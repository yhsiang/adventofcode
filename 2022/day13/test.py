def compare(one, two, enum=0):
    lst = [one, two]
    while enum < (length := min(len(one), len(two))):
        val_one, val_two = one[enum], two[enum]
        if val_one == val_two:
            enum += 1
            if enum == length:
                return len(one) < len(two)
        else:
            type_one, type_two = type(val_one), type(val_two)
            if type_one != type_two:
                if "]" in [val_one, val_two]:
                    return val_one == "]"
                to_edit = lst[next((i for i,x in enumerate([type_one, type_two]) if x == int))]
                to_edit.insert(enum + 1, "]")
                to_edit.insert(enum, "[")
            else:
                return val_one < val_two if type_one == int else val_one == "]" and val_two == "["

with open("input", "r") as file:
    data = file.read().replace("[", "[,").replace("]", ",]")
    data = [x.splitlines() for x in data.split('\n\n')]
    p1 = 0
    for e, packets in enumerate(data):
        for i, packet in enumerate(packets):
            data[e][i] = [int(x) if x.isdigit() else x for x in packet.split(",") if x]
        if compare(data[e][0][:], data[e][1][:]):
            p1 += e + 1
    data = sum(data, []) + [(two_extra := ["[", "[", 2, "]", "]"]), (six_extra := ["[", "[", 6, "]", "]"])]
    sorted_data = [data[0]]
    for x in data[1:]:
        for e, y in enumerate(sorted_data):
            if compare(x, y):
                sorted_data.insert(e, x)
                break
        else:
            sorted_data.append(x)
    print("day 13: ", p1, (sorted_data.index(two_extra) + 1) * (sorted_data.index(six_extra) + 1))