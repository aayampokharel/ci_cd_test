shopt -s globstar
while read r;do
echo $r
if [[ $r = scripts/** ]];then
echo " yes this is the one " $r
fi
done < ./output.txt
