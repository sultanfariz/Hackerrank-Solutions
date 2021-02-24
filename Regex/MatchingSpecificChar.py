Regex_Pattern = r'^[1-3][1,2,0][x,s,0][3,0,A,a][x,s,u][\.\,]$'	# Do not delete 'r'.

import re

print(str(bool(re.search(Regex_Pattern, input()))).lower())
