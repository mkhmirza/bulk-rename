#!/bin/bash 

# create number given number of files of specified extension, and a random name

# how to many files to create
TIMES=$1
# echo "Running file maker" $TIMES "times"
# get the specified extension if any
ext=".${2}"


# print out the information passed as arguments
echo "Creating " $TIMES " Files"
if [ $ext == "." ]; then
    echo "Extension: No Extension Given"
else
    echo "Extension: " $ext
fi

echo
# check if a test/ folder exists
if [ -d "./test/" ] 
then
    echo "Directory does exists." 
else
    # create a directory ./test/
    echo "Error: Directory test/ does not exists."
    mkdir test/
    echo "Directory test/ created."
    echo
fi



for (( c=0; c<=$TIMES; c++ ))
do
    # create a random string of a file, having characters a-f 0-9 and of 5 len
    name=$(cat /dev/urandom | tr -cd 'a-f0-9' | head -c 10)

    # check to see if a extension is given or not
    if [[ $ext == "." ]]; then
        # concat randomly created file name without a given extension i.e filname
        filename="${name}"
    else 
        # concat randomly created file name with a given extension i.e filname.txt
        filename="${name}${ext}"
    fi

    echo "Creating: test/"$filename
    # create the file in the test folder
    touch test/$filename
    echo "File Created: test/"$filename 
    echo
done
