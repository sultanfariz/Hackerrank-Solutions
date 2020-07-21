def swap_case(s):
    x=''
    for _ in s:
        if(_.isupper()): x+=_.lower()
        elif(_.islower()): x+=_.upper()
        else: x+=_
    return x

if __name__ == '__main__':
    s = input()
    result = swap_case(s)
    print(result)
