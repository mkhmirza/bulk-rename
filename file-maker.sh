#!/bin/bash 

# create number given number of files of specified extension, and a random name

# how to many files to create
TIMES=$1
echo "Running file maker" $TIMES "times"
echo 
# get the specified extension if any
ext=$2

# check if a folder exists
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
    # concat randomly created file name with the given extension i.e filname.ext
    filename="${name}.${ext}"
    echo "Creating: test/"$filename
    # create the file in the test folder
    touch test/$filename
    echo "File Created: test/"$filename 
    echo
done
