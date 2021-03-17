Start_Regex_Pattern = r'^hackerrank'
End_Regex_Pattern = r'hackerrank$'

import re

for _ in range(int(input())):
    Test_String = input()
    match1 = re.findall(Start_Regex_Pattern, Test_String)
    match2 = re.findall(End_Regex_Pattern, Test_String)
    if len(match1) == 0 and len(match2) == 0:
        print(-1)
    elif len(match1) != 0 and len(match2) != 0:
        print(0)
    elif len(match1) != 0 and len(match2) == 0:
        print(1)
    elif len(match1) == 0 and len(match2) != 0:
        print(2)
