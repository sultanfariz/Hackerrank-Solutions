import re

n=0
for _ in range(int(input())):
    match = re.search(r'hackerrank', input(), re.IGNORECASE)
    if match:
        n+=1
print(n)
