
if [[ $(ls -d $1 2> /dev/null) = $1 ]]; then
    echo "directory exists"
else 
echo "directory doesnot exist"

fi
