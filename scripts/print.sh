echo " enter filename"
read fileName         
echo "the file is $fileName and the wc is : $(wc -l < $fileName)"
