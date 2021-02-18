#!/bin/sh
for x in `cat`; do 
    array=( "${array[@]}" $x )
done
array=("${array[@]}" "${array[@]}" "${array[@]}")
echo "${array[@]}"
