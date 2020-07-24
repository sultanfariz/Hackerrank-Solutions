if __name__ == '__main__':
    s = input()
    cmd=[str.isalnum, str.isalpha, str.isdigit, str.islower, str.isupper]
    for i in cmd:
        print(any(i(_) for _ in s))
