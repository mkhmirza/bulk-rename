#!/bin/bash 

# create number given number of files of specified extension, and a random name
# usage: ./file-maker.sh $1 $2 $3
# $1: number of files and folders
# $2: extension of the create file, leave empty for files without extension
# $3: create a folder or a file

# how to many files to create
TIMES=$1
# create files or folder
mode=$2
# echo "Running file maker" $TIMES "times"
# get the specified extension if any
ext=".${3}"


# if mode is empty then set mode as file
# meaning create file
if [ $mode == "" ]; then 
    echo "No mode given. Setting mode to file"
    $mode="file"
fi

# print out the information passed as arguments
echo "Creating " $TIMES " Files"
if [ $ext == "." ]; then
    echo "Extension: No Extension Given"
else
    echo "Extension: " $ext
fi
echo "Mode: " $mode

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

# create a random files/folders in test/
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

    # check the mode of opertion
    if [[ $mode == "file" ]]; then
        echo "Creating: test/"$filename
        touch test/$filename
        echo "File Created: test/"$filename 
    # if mode is given as folder
    else
        foldername="${name}"
        echo "Creating: test/"$foldername
        mkdir -p test/$foldername
        echo "Folder Created: test/"$foldername
    fi 
    echo
done
