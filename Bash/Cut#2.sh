#!/bin/sh
while read x
do echo $x | cut -c 2,7
done
