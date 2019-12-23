arr = []
trash = [9, 7, 5, 3, 1]


def swap(a, b):
    arr.append([a, b])


def compare(a, b):
    return trash[a] > trash[b]


for i in range(5):
    for j in range(5-i-1):
        if compare(j,j+1):
            swap(j,j+1)

print(arr)