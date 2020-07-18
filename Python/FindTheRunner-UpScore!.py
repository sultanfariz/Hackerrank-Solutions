if __name__ == '__main__':
    n = int(input())
    arr = list(map(int, input().split()))
    s=max(arr); i=0
    while i<n:
        if max(arr)==s: arr.remove(max(arr))
        i+=1
    print(max(arr))
