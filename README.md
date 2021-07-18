# Bulk Renamer
Rename multiple files using terminal, based on a pattern. See [Note](#note) before using to rename your personal files/folders.

## Installing Go Binary
To create/install a go binary use
```
go install bulk-renamer.go
```
The go binary are stored in `bin/` dir. To setup a Golang workspace see [how to setup a golang workspace](https://golang.org/doc/gopath_code#Workspaces).

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

### File Maker
File Maker is a bash script to create files or folders to test **bulk-renamer**. 
```
file-maker.sh <no-of-files-create> <mode> <extension-of-files>
```
For creating folders use
```
file-maker.sh 5 folder 
```
For creating files use
```
file-maker.sh 5 file txt
```
See [file-maker.sh](file-maker.sh) for more understanding.
## Note
Use this program with caution, as it can overwrite files having same name. Make sure you are passing the right arguments while renaming them.