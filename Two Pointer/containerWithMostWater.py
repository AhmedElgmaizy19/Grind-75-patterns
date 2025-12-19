
def max_area(height):
    start = 0
    end = len(height) - 1
    maxArea = 0
    
    while start < end :
        width = end - start
        height= int(min(height[start],height[end]))
        area = width * height
        if area > maxArea:
            maxArea = area
        
        if height[start] < height[end]:
            start+=1
        else :
            end-=1
    return maxArea