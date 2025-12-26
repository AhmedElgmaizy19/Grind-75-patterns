

def three_sum_value(nums,target):
    nums.sort()
    
    for i in range(len(nums)-2):
        low=i+1
        high=len(nums)-1
        while low < high:
            res = nums[i] + nums[low] + nums[high]
            if res == target:
                return True
            elif res < target:
                low+=1
            else:
                high-=1
    return False
            