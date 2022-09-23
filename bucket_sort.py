def bucketSort(array):
    maxValue = max(array)
    minValue = min(array)
    n = len(array)

    numberOfBucket = (maxValue - minValue)//n + 1

    # create empty buckets
    buckets = []
    for i in range(numberOfBucket):
        buckets.append([])
    
    # allocate elements to their corresponding bucket
    for i in range(len(array)):
        element = array[i]
        bucketIdx = (element - minValue)//n
        buckets[bucketIdx].append(element)

    print("unsorted buckets: ", buckets)

    # sort individual buckets using insertion sort
    for i in range(numberOfBucket):
        insertionSort(buckets[i])
    
    print("sorted buckets: ", buckets)

    # combine sorted buckets to a sorted array
    sortedIdx = 0
    for i in range(numberOfBucket):
        currentBucket = buckets[i]
        for j in range(len(currentBucket)):
            array[sortedIdx] = currentBucket[j]
            sortedIdx += 1

    return array

def insertionSort(array):
    for i in range(1, len(array)):
        j = i
        while j > 0 and array[j-1] > array[j]:
            swap(j-1, j, array)
            j -= 1

def swap(i, j, array):
    array[i], array[j] = array[j], array[i]

if __name__ == "__main__":
    arrays = [[15, 5, 2, 9, 8, 4, 22], [15, 8, 4, -9, 5, -2, 22]]
    for array in arrays:
        print("unsorted array", array)
        print("sorted array", bucketSort(array))
        print("\n")

"""
output:
('unsorted array', [15, 5, 2, 9, 8, 4, 22])
('unsorted buckets: ', [[5, 2, 8, 4], [15, 9], [22]])
('sorted buckets: ', [[2, 4, 5, 8], [9, 15], [22]])
('sorted array', [2, 4, 5, 8, 9, 15, 22])


('unsorted array', [15, 8, 4, -9, 5, -2, 22])
('unsorted buckets: ', [[-9], [4, -2], [8, 5], [15], [22]])
('sorted buckets: ', [[-9], [-2, 4], [5, 8], [15], [22]])
('sorted array', [-9, -2, 4, 5, 8, 15, 22])

"""