# Bulk Renamer
Rename multiple files using terminal, based on a pattern. See [Note](##Note) before using to rename your personal files/folders.

## Installing Go Binary
To create/install a go binary use
```
go install bulk-renamer.go
```
The go binary are stored in `pkg/` dir. To setup a Golang workspace see [how to setup a golang workspace](https://golang.org/doc/gopath_code#Workspaces).

## Usage
using **go**:
```
cd bulk-renamer
go build bulk-renamer.go
```
After building the binary run the following command for help
```
bulk-renamer -h
```
For Renaming files use
```
bulk-renamer -path <path-to-folder> -pattern <pattern> 
```
Make sure to add the trailing `/` in the folder path for example:<br /> 
```
path/to/some/folder/ # right way
path/to/some/folder  # wrong way
```

### Renaming only folders within a DIR
To rename folders within a directory containing other files use 
```
bulk-renamer <path-to-folder> -pattern <pattern> -f
```
`-f` flag makes sure that only folders within a directory are renamed.

### Renaming Files With A Specific Extension
To rename all files of a specific extension, run
```
bulk-renamer -path <path-to-folder> -pattern <pattern>  -extension <file-extension> 
```

### Starting Point for the Counter
Starting number for the counter can be changed as
```
bulk-renamer -path <path-to-folder> -starting-point <starting-number> 
```

### Dry Run
A dry run can be performed on a **test/** folder. To create multiple files in a one go run 
```
./file-maker.sh <no-of-files-create> <extension-of-files>
```
for example
```
./file-maker.sh 5 txt
```
creates a folder with a name **test** in my current directory and create `5` file with random strings and `txt` as file extension. The output is given below.
```
Running file maker 5 times

Directory does exists.
Creating: test/6f2d828cd7.txt
File Created: test/6f2d828cd7.txt

Creating: test/6fff621041.txt
File Created: test/6fff621041.txt

Creating: test/c584275ec5.txt
File Created: test/c584275ec5.txt

Creating: test/d82e655d1f.txt
File Created: test/d82e655d1f.txt

Creating: test/8239fee994.txt
File Created: test/8239fee994.txt

Creating: test/4e6906fbeb.txt
File Created: test/4e6906fbeb.txt
```
After creating the files using your prefered method, run the go program 
```
bulk-renamer -pattern <pattern> -extension <file-extension> -dry-run
```
This will rename all the files in `test` folder having extension `txt`.<br ><br >
**Note**: Run ./file-maker.sh before running `bulk-renamer -pattern <pattern> -extension <file-extension> -dry-run` option.

## Note
Use this program with caution, as it can overwrite files having same name. Make sure you are passing the right arguments while renaming them. It is best to run `-dry-run` option just to see how the program works, then running it on personal files.