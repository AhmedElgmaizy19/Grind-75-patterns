

def sort_colors(colors):
    start , end , current = 0 , len(colors) - 1 , 0
    while current <= end:
        match colors[current]:
            case 0:
                if colors[start] != 0:
                    colors[start] , colors[current] = colors[current] , colors[start]
                start += 1
                current += 1
            case 1:
                current += 1
            
            case 2:
                if colors[end] != 2:
                    colors[end] , colors[current] = colors[current] , colors[end]
                end -= 1
    return colors