#!/bin/sh
i=0
while read x
do array[$i]=$x
((i=i+1))
done

echo ${array[3]}
