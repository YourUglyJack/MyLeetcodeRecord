def n_length_combo(list, n):
    if n == 0:
        return [[]]
    l = []
    for i in range(0, len(list)):
        m = list[i]
        remainlist = list[i + 1:]

        remainlist_combo = n_length_combo(remainlist, n - 1)
        for p in remainlist_combo:
            l.append([m, *p])
    return l


arr = "abc"
list = [x for x in arr]
res = n_length_combo(list, 2)
