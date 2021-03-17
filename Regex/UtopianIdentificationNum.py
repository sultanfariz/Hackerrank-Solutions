Regex_Pattern = r"^[a-z]{0,3}\d{2,8}[A-Z]{3,}$"    # Do not delete 'r'.

import re

for _ in range(int(input())):
    Test_String = input()
    match = re.findall(Regex_Pattern, Test_String)
    # print("Number of matches :", len(match))
    if len(match) == 0:
        print("INVALID")
    else : print("VALID")
