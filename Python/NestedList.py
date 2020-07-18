if __name__ == '__main__':
    ar=list(); arr=list(); s=100; f=100
    for _ in range(int(input())):
        name = input()
        score = float(input())
        arr.append([name,score])
        if score<f: s=f; f=score
        elif score>f and score<s: s=score 
    for _ in sorted(i[0] for i in arr if i[1]==s): print(_)
