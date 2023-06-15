# to read value $1 from file $2
echo $(sed -n 's/^'$1'=\(.*$\)/\1/p' $2)