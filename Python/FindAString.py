def count_substring(string, sub_string):
    n = 0
    a = len(string); b = len(sub_string)
    for _ in range(a-b+1):
        if(string[0+_ : b+_] == sub_string):
            n+=1
    return n

if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()
    
    count = count_substring(string, sub_string)
    print(count)
