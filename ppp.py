def getMaximumGrossValue(arr):
    if not arr:  # If the list is empty, return 0
        return 0

    max_so_far = arr[0]
    current_max = arr[0]
    
    for i in range(1, len(arr)):
        current_max = max(arr[i], current_max + arr[i])
        max_so_far = max(max_so_far, current_max)
        
    return max_so_far


# Example usage:
sample_arr = [4,-8,2,-10,3,-20]
getMaximumGrossValue(sample_arr)

print(getMaximumGrossValue(sample_arr))
