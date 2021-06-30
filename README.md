# Bulk Renamer
Rename multiple files using terminal, based on a pattern.

## Usage
using **go**:
```
cd rename-files
go build rename.go
```
After building the binary run the following command for help
```
./rename -h
```
For Renaming files use
```
./rename -folder <path-to-folder> -pattern <pattern> 
```
Make sure to add the trailing `/` in the folder path for example:<br /> 
```
path/to/some/folder/ # right way
path/to/some/folder  # wrong way
```

### Renaming only folders within a DIR
To rename folders within a directory containing other files use 
```
./rename <path-to-folder> -pattern <pattern> -f
```
`-f` flag makes sure that only folders within a directory are renamed.

### Renaming Files With A Specific Extension
To rename all files of a specific extension, run
```
./rename -folder <path-to-folder> -pattern <pattern>  -extension <file-extension> 
```