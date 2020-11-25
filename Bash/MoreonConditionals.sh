read x
read y
read z

if [ $x = $y ] && [ $x = $z ] && [ $y = $z ]
then echo "EQUILATERAL"
elif [ $x != $y ] && [ $x != $z ] && [ $y != $z ]
then echo "SCALENE"
else echo "ISOSCELES"
fi
