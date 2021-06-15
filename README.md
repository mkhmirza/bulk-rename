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
