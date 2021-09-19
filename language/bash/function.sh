#!/bin/bash

hello () {
	echo hello $1 and $2	
	return 0
}

hello1 () {
	echo hello $1 and $2
	return 1
}

# hello joe jack
# hello1 joe jac

# echo "The previous function has a return value of $?"

# substitution
lines_in_file () {
	cat $1 | wc -l
	local file="./file"
}

no_line=$(lines_in_file $1)

echo "The file $1 has $no_line lines in it" 
file="./file"

if [ $file ]; then 
	echo "the $file is there"
elif [[ ! $file ]]; then
	echo "noo"
fi