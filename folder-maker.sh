#!/bin/bash 

# create number given number of folder with a random name

# how to many folder to create
TIMES=$1
echo "Running file maker" $TIMES "times"
echo 

# check if a folder exists
if [ -d "./test/" ] 
then
    echo "Directory does exists." 
else
    # create a directory ./test/
    echo "Error: Directory test/ does not exists. Creating.."
    mkdir test/
    echo "Directory test/ created."
    echo
fi


for (( c=0; c<=$TIMES; c++ ))
do
    # create a random string of a folder, having characters a-f 0-9 and of 5 len
    name=$(cat /dev/urandom | tr -cd 'a-f0-9' | head -c 10)
    # concat randomly created folder name with the given extension i.e filname.ext
    foldername="${name}"
    echo "Creating: test/"$foldername
    # create the folder in the test folder
    mkdir -p test/$foldername
    echo "Folder Created: test/"$filename 
    echo
done
