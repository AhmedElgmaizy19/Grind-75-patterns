
# naive solution O(n^2) (Not optimal)
# def productExceptSelf(nums):
#     result = []
#     for i in range(len(nums)):
#         pro = 1
#         for j in range(len(nums)):
#             if i != j:
#                 pro *= nums[j]
#         result.append(pro)
#     return result


def productExceptSelf(nums):
    n = len(nums)
    result = [1]*n
    start = 0
    end = n-1
    left_product = 1
    right_product = 1
    while start < n and end >= 0:
      result[start]*=left_product
      result[end]*=right_product
      left_product*=nums[start]
      right_product*=nums[end]
      start+=1
      end-=1
    return result
      
x= [-3,3,2,-1,5]
print(productExceptSelf(x))