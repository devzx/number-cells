x = [1, 2, 50, 3213, 12, 2, 1]


def cell(a, c):
    cell_width = len(str(max(a)))
    row_art = '-' * cell_width
    padding = '{:>' + str(cell_width) + '}|'

    print("+", end="")
    column_count = int()
    for _ in a:
        if column_count == c:
            break
        column_count += 1
        print(row_art + "+", end="")

    print("\n|", end="")
    column_count = 0
    for i in a:
        if column_count == c:
            print("\n+", end="")
            for _ in range(c):
                print(row_art + "+", end="")
            print("\n|", end="")
            column_count = 0
        column_count += 1
        print(padding.format(i), end="")
    print("\n+", end="")
    if len(a) % c == 0 and c != 1:
        for _ in range(c):
            print(row_art + "+", end="")
    elif c == 1:
        print(row_art + "+")
    else:
        for _ in range(len(a) % c):
            print(row_art + "+", end="")


cell(x, 4)
