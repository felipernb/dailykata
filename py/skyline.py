#problem specification: http://acm.uva.es/p/v1/105.html

buildings = [(1,11,5), (2,6,7), (3,13,9), (12,7,16), (14,3,25), (19,18,22), (23,13,29), (24,4,28)]

def skyline(buildings):
    #initialize the skyline, setting 0 in all points
    line = [0 for i in xrange(max(b[2] for b in buildings)+1)]
    for b in buildings:
        for x in xrange(b[0],b[2]):
            line[x] = max(line[x], b[1])
    skyline = []
    actual_height = 0
    for (x, y) in enumerate(line):
        if y != actual_height:
            skyline.append((x,y))
            actual_height = y

    return skyline

print skyline(buildings)
