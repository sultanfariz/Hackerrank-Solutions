read -n1 x
if [ "$x" = "y" ] || [ "$x" = "Y" ]
then echo "YES"
# elif [ "$x" = "Y" ]
# then echo "YES"
else echo "NO"
fi
