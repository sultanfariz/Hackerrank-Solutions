#!/bin/sh
while read x
do echo $x|cut -d ' ' -f1-3
done
